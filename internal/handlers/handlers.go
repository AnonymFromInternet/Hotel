package handlers

import (
	"github.com/anonymfrominternet/Hotel/internal/config"
	"github.com/anonymfrominternet/Hotel/internal/models"
	"github.com/anonymfrominternet/Hotel/internal/render"
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

// MainPage is a GET handler for the main page
func (repo *Repository) MainPage(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, "main.page.tmpl", &models.TemplateData{})
}

// AboutPage is a GET handler for the about page
func (repo *Repository) AboutPage(writer http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["defaultData"] = "default value"
	render.Template(writer, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Generals is a GET handler for the generals page
func (repo *Repository) Generals(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, "generals.page.tmpl", &models.TemplateData{})
}

// President is a GET handler for the president page
func (repo *Repository) President(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, "president.page.tmpl", &models.TemplateData{})
}

// Availability is a GET handler for the search-availability page
func (repo *Repository) Availability(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, "search-availability.page.tmpl", &models.TemplateData{})
}

// Contact is a GET handler for the Contact page
func (repo *Repository) Contact(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, "contact.page.tmpl", &models.TemplateData{})
}

// Reservation is a GET handler for the reservation page
func (repo *Repository) Reservation(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, "reservation.page.tmpl", &models.TemplateData{})
}
