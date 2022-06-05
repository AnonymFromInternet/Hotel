package handlers

import (
	"encoding/json"
	"github.com/anonymfrominternet/Hotel/internal/config"
	"github.com/anonymfrominternet/Hotel/internal/driver"
	"github.com/anonymfrominternet/Hotel/internal/forms"
	"github.com/anonymfrominternet/Hotel/internal/helpers"
	"github.com/anonymfrominternet/Hotel/internal/models"
	"github.com/anonymfrominternet/Hotel/internal/render"
	"github.com/anonymfrominternet/Hotel/internal/repository"
	"github.com/anonymfrominternet/Hotel/internal/repository/dbRepo"
	"net/http"
	"strconv"
	"time"
)

type Repository struct {
	AppConfig *config.AppConfig
	DB        repository.DatabaseRepository
}

// Repo is the Repository for the handlers
var Repo *Repository

// NewRepo gets appConfig from main()
func NewRepo(appConfig *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		AppConfig: appConfig,
		DB:        dbRepo.NewPostgresDBRepo(appConfig, db.SQL),
	}
}

// NewHandlers sets value for the var Repo
func NewHandlers(repo *Repository) {
	Repo = repo
}

// GET HANDLERS

// MainPage is a GET handler for the main page
func (repo *Repository) MainPage(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, request, "main.page.tmpl", &models.TemplateData{})
}

// AboutPage is a GET handler for the about page
func (repo *Repository) AboutPage(writer http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["defaultData"] = "default value"
	render.Template(writer, request, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Generals is a GET handler for the generals page
func (repo *Repository) Generals(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, request, "generals.page.tmpl", &models.TemplateData{})
}

// President is a GET handler for the president page
func (repo *Repository) President(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, request, "president.page.tmpl", &models.TemplateData{})
}

// Availability is a GET handler for the search-availability page
func (repo *Repository) Availability(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, request, "search-availability.page.tmpl", &models.TemplateData{})
}

// Contact is a GET handler for the Contact page
func (repo *Repository) Contact(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, request, "contact.page.tmpl", &models.TemplateData{})
}

// Reservation is a GET handler for the reservation page
func (repo *Repository) Reservation(writer http.ResponseWriter, request *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservationPageInputs"] = emptyReservation

	render.Template(writer, request, "reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON is a GET handler. This handler sends back JSON data about availability
func (repo *Repository) AvailabilityJSON(writer http.ResponseWriter, request *http.Request) {
	response := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(response, "", "   ")
	if err != nil {
		helpers.ServerError(writer, err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write(out)
}

// ReservationSummary is a GET handler for the reservation-summary page
func (repo *Repository) ReservationSummary(writer http.ResponseWriter, request *http.Request) {
	// Getting the data from the request context and putting it to AppConfig.Session and trying to type assertion
	// This data was added in the PostReservation handler
	reservationPageInputs, ok := repo.AppConfig.Session.Get(request.Context(),
		"reservationPageInputs").(models.Reservation)
	if !ok {
		repo.AppConfig.ErrorLog.Println("Cannot get data from reservation")
		repo.AppConfig.Session.Put(request.Context(), "Error", "Cannot get data from reservation")
		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
		return
	}

	repo.AppConfig.Session.Remove(request.Context(), "reservationPageInputs")
	data := make(map[string]interface{})
	data["reservationPageInputs"] = reservationPageInputs

	render.Template(writer, request, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

// GET HANDLERS

// POST HANDLERS

// PostAvailability is a POST handler for the search-availability page
func (repo *Repository) PostAvailability(writer http.ResponseWriter, request *http.Request) {
	// Getting data from form by the POST method
	startDateInString := request.Form.Get("start_date")
	endDateInString := request.Form.Get("end_date")
	// Getting data from form by the POST method

	datesLayout := "2006-01-02"
	startDate, err := time.Parse(datesLayout, startDateInString)
	if err != nil {
		helpers.ServerError(writer, err)
		return
	}
	endDate, err := time.Parse(datesLayout, endDateInString)
	if err != nil {
		helpers.ServerError(writer, err)
		return
	}

	rooms, err := repo.DB.AllAvailableRooms(startDate, endDate)
	if err != nil {
		helpers.ServerError(writer, err)
		return
	}

	if len(rooms) == 0 {
		repo.AppConfig.Session.Put(request.Context(), "Error", "no available rooms")
		http.Redirect(writer, request, "/search-availability", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["rooms"] = rooms

	reservation := models.Reservation{
		StartDate: startDate,
		EndDate:   endDate,
	}
	repo.AppConfig.Session.Put(request.Context(), "reservation", reservation)

	render.Template(writer, request, "choose-room.page.tmpl", &models.TemplateData{Data: data})
}

// PostReservation is a POST handler for the reservation page
func (repo *Repository) PostReservation(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		helpers.ServerError(writer, err)
		return
	}

	// Getting data which was added by user in inputs
	startDateAsString := request.Form.Get("start_date")
	endDateAsString := request.Form.Get("end_date")

	datesLayout := "2006-01-02"

	startDate, err := time.Parse(datesLayout, startDateAsString)
	if err != nil {
		helpers.ServerError(writer, err)
		return
	}
	endDate, err := time.Parse(datesLayout, endDateAsString)
	if err != nil {
		helpers.ServerError(writer, err)
		return
	}

	roomId, err := strconv.Atoi(request.Form.Get("room_id"))
	if err != nil {
		helpers.ServerError(writer, err)
		return
	}

	reservationPageInputs := models.Reservation{
		FirstName: request.Form.Get("first_name"),
		LastName:  request.Form.Get("last_name"),
		Email:     request.Form.Get("email"),
		Phone:     request.Form.Get("phone"),
		StartDate: startDate,
		EndDate:   endDate,
		RoomId:    roomId,
	}
	// Getting data which was added by user in inputs

	form := forms.New(request.PostForm)
	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservationPageInputs"] = reservationPageInputs

		render.Template(writer, request, "reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	// Adding info to db, to the Reservations Table, and getting id of new added item from this table
	reservationId, err := repo.DB.InsertReservation(reservationPageInputs)
	if err != nil {
		helpers.ServerError(writer, err)
		return
	}
	// Adding info to db

	// Creating RoomRestriction and adding this data to the db, in the RoomRestrictions table
	restriction := models.RoomRestriction{
		StartDate:     startDate,
		EndDate:       endDate,
		RoomId:        roomId,
		ReservationId: reservationId,
		RestrictionId: 1,
	}
	err = repo.DB.InsertRoomRestriction(restriction)
	if err != nil {
		helpers.ServerError(writer, err)
		return
	}

	repo.AppConfig.Session.Put(request.Context(), "reservationPageInputs", reservationPageInputs)
	http.Redirect(writer, request, "/reservation-summary", http.StatusSeeOther)

}

// POST HANDLERS
