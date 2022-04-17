package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/anonymfrominternet/Hotel/internal/config"
	"github.com/anonymfrominternet/Hotel/internal/driver"
	"github.com/anonymfrominternet/Hotel/internal/forms"
	"github.com/anonymfrominternet/Hotel/internal/helpers"
	"github.com/anonymfrominternet/Hotel/internal/models"
	"github.com/anonymfrominternet/Hotel/internal/render"
	"github.com/anonymfrominternet/Hotel/internal/repository"
	dbrepo "github.com/anonymfrominternet/Hotel/internal/repository/dbRepo"
	"net/http"
	"strconv"
	"time"
)

// Repo Repository pattern
var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

func NewTemplates(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}
func NewHandlers(r *Repository) {
	Repo = r
}

// MainPage is the handler for GET requests on the main page
func (r *Repository) MainPage(writer http.ResponseWriter, request *http.Request) {

	render.Template(writer, request, "main.page.gohtml", &models.TemplateData{})
}

// AboutPage is the handler for GET requests on the about page
func (r *Repository) AboutPage(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, request, "about.page.gohtml", &models.TemplateData{})
}

// PresidentPage is the handler for GET requests on the PresidentPage
func (r *Repository) PresidentPage(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, request, "president.page.gohtml", &models.TemplateData{})
}

// BusinessPage is the handler for GET requests on the BusinessPage
func (r *Repository) BusinessPage(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, request, "business.page.gohtml", &models.TemplateData{})
}

// Calendar is the handler for GET requests on the Calendar page
func (r *Repository) Calendar(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, request, "calender.page.gohtml", &models.TemplateData{})
}

// PostCalendar is the handler for POST requests on the Calendar page
func (r *Repository) PostCalendar(writer http.ResponseWriter, request *http.Request) {
	// Delete?
	start := request.Form.Get("start")
	end := request.Form.Get("end")
	_, err := writer.Write([]byte(fmt.Sprintf("Starting date is %s and ending date is %s", start, end)))
	if err != nil {
		fmt.Println("Error in handlers / PostCalendar / writer.Write")
	}
	// Delete?
}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

// CalendarJSON reforms data from JSON to instance of jsonResponse struct
func (r *Repository) CalendarJSON(writer http.ResponseWriter, request *http.Request) {
	resp := jsonResponse{true, "Available"}

	out, err := json.MarshalIndent(resp, "", "   ")
	if err != nil {
		helpers.ServerError(writer, err)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	_, err = writer.Write(out)
	if err != nil {
		helpers.ServerError(writer, err)
		return
	}
}

// Contacts is the handler for POST requests on the Contacts page
func (r *Repository) Contacts(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, request, "contacts.page.gohtml", &models.TemplateData{})
}

// Reservation is the handler for POST requests on the Reservation page
func (r *Repository) Reservation(writer http.ResponseWriter, request *http.Request) {
	// Why? When a POST method is called, and errors are not empty, then it works without it. Because personalData in PostReservation() parses all data
	var emptyPersonalData models.Reservation
	data := make(map[string]interface{})
	data["personalData"] = emptyPersonalData
	// ?

	render.Template(writer, request, "personal-data.page.gohtml", &models.TemplateData{
		Form: forms.NewForm(nil),
		Data: data,
	})
}

func (r *Repository) PostReservation(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		helpers.ServerError(writer, err)
		return
	}

	// Reformatting data string parsed from the page to the layout format:
	layout := "2006-01-03"
	startDate, err := time.Parse(layout, request.Form.Get("start_date"))
	if err != nil {
		helpers.ServerError(writer, err)
	}

	endDate, err := time.Parse(layout, request.Form.Get("end_date"))
	if err != nil {
		helpers.ServerError(writer, err)
	}

	roomId, err := strconv.Atoi(request.Form.Get("room_id"))
	if err != nil {
		helpers.ServerError(writer, err)
	}

	reservation := models.Reservation{
		FirstName: request.Form.Get("first_name"),
		LastName:  request.Form.Get("last_name"),
		Email:     request.Form.Get("email"),
		Phone:     request.Form.Get("phone"),
		StartDate: startDate,
		EndDate:   endDate,
		RoomId:    roomId,
	}

	form := forms.NewForm(request.PostForm)

	//form.Has("first_name", request)
	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.Template(writer, request, "personal-data.page.gohtml", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	// Puts data into Reservations Table in the database:
	err = r.DB.InsertReservation(reservation)
	if err != nil {
		helpers.ServerError(writer, err)
	}
	// Puts data into Reservations Table in the database

	r.App.Session.Put(request.Context(), "reservation", reservation)

	http.Redirect(writer, request, "/after-personal-data", http.StatusSeeOther)
}

func (r *Repository) ReservationSummary(writer http.ResponseWriter, request *http.Request) {
	personalData, ok := r.App.Session.Get(request.Context(), "personalData").(models.Reservation)
	if !ok {
		r.App.ErrorLog.Println("Cannot get error from session")
		r.App.Session.Put(request.Context(), "error", "Cannot get personalData from session")
		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
		return
	}
	r.App.Session.Remove(request.Context(), "personalData")

	data := make(map[string]interface{})
	data["personalData"] = personalData
	render.Template(writer, request, "after-personal-data.page.gohtml", &models.TemplateData{
		Data: data,
	})
}
