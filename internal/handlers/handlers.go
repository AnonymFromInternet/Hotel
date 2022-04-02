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
)

// Repo Repository pattern
var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}
func NewHandlers(r *Repository) {
	Repo = r
}

func (r *Repository) MainPage(writer http.ResponseWriter, request *http.Request) {

	render.RenderTemplate(writer, request, "main.page.gohtml", &models.TemplateData{})
}
func (r *Repository) AboutPage(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "about.page.gohtml", &models.TemplateData{})
}

func (r *Repository) PresidentPage(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "president.page.gohtml", &models.TemplateData{})
}
func (r *Repository) BusinessPage(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "business.page.gohtml", &models.TemplateData{})
}

func (r *Repository) Calender(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "calender.page.gohtml", &models.TemplateData{})
}

func (r *Repository) PostCalender(writer http.ResponseWriter, request *http.Request) {
	// Delete?
	start := request.Form.Get("start")
	end := request.Form.Get("end")
	_, err := writer.Write([]byte(fmt.Sprintf("Starting date is %s and ending date is %s", start, end)))
	if err != nil {
		fmt.Println("Error in handlers / PostCalender / writer.Write")
	}
	// Delete?
}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (r *Repository) CalenderJSON(writer http.ResponseWriter, request *http.Request) {
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

func (r *Repository) Contacts(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "contacts.page.gohtml", &models.TemplateData{})
}

func (r *Repository) PersonalData(writer http.ResponseWriter, request *http.Request) {
	// Why? When a POST method is called, and errors are not empty, then it works without it. Because personalData in PostPersonalData() parses all data
	var emptyPersonalData models.PersonalData
	data := make(map[string]interface{})
	data["personalData"] = emptyPersonalData
	// ?

	render.RenderTemplate(writer, request, "personal-data.page.gohtml", &models.TemplateData{
		Form: forms.NewForm(nil),
		Data: data,
	})
}

func (r *Repository) PostPersonalData(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		helpers.ServerError(writer, err)
		return
	}

	personalData := models.PersonalData{
		FirstName: request.Form.Get("first_name"),
		LastName:  request.Form.Get("last_name"),
		Email:     request.Form.Get("email"),
		Phone:     request.Form.Get("phone"),
	}

	form := forms.NewForm(request.PostForm)

	//form.Has("first_name", request)
	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["personalData"] = personalData

		render.RenderTemplate(writer, request, "personal-data.page.gohtml", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
	r.App.Session.Put(request.Context(), "personalData", personalData)

	http.Redirect(writer, request, "/after-personal-data", http.StatusSeeOther)
}

func (r *Repository) AfterPersonalData(writer http.ResponseWriter, request *http.Request) {
	personalData, ok := r.App.Session.Get(request.Context(), "personalData").(models.PersonalData)
	if !ok {
		r.App.ErrorLog.Println("Cannot get error from session")
		r.App.Session.Put(request.Context(), "error", "Cannot get personalData from session")
		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
		return
	}
	r.App.Session.Remove(request.Context(), "personalData")

	data := make(map[string]interface{})
	data["personalData"] = personalData
	render.RenderTemplate(writer, request, "after-personal-data.page.gohtml", &models.TemplateData{
		Data: data,
	})
}
