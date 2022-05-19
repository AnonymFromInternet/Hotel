package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/anonymfrominternet/Hotel/internal/config"
	"github.com/anonymfrominternet/Hotel/internal/forms"
	"github.com/anonymfrominternet/Hotel/internal/models"
	"github.com/anonymfrominternet/Hotel/internal/render"
	"log"
	"net/http"
)

type Repository struct {
	AppConfig *config.AppConfig
}

// Repo is the Repository for the handlers
var Repo *Repository

// NewRepo gets appConfig from main()
func NewRepo(appConfig *config.AppConfig) *Repository {
	return &Repository{
		AppConfig: appConfig,
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
	var emptyReservation models.ReservationPageInputtedData
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
		fmt.Println("cannot convert response to JSON")
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, _ = writer.Write(out)
}

func (repo *Repository) ReservationSummary(writer http.ResponseWriter, request *http.Request) {
	reservationPageInputs, ok := repo.AppConfig.Session.Get(request.Context(),
		"reservationPageInputs").(models.ReservationPageInputtedData)
	if !ok {
		log.Print("Cannot assert data type")
		return
	}
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
	start := request.Form.Get("start_date")
	end := request.Form.Get("end_date")
	// Getting data from form by the POST method

	_, _ = writer.Write([]byte(fmt.Sprintf("Start is %s, end is %s", start, end)))

}

// PostReservation is a POST handler for the reservation page
func (repo *Repository) PostReservation(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		fmt.Println("cannot parse data from request", err)
		return
	}

	// Getting data which was added by user in inputs
	reservationPageInputs := models.ReservationPageInputtedData{
		FirstName: request.Form.Get("first_name"),
		LastName:  request.Form.Get("last_name"),
		Email:     request.Form.Get("email"),
		Phone:     request.Form.Get("phone"),
	}
	// Getting data which was added by user in inputs

	form := forms.New(request.PostForm)
	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 3, request)
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

	repo.AppConfig.Session.Put(request.Context(), "reservationPageInputs", reservationPageInputs)
	http.Redirect(writer, request, "/reservation-summary", http.StatusSeeOther)

}

// POST HANDLERS
