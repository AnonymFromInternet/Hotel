package handlers

import (
	"github.com/anonymfrominternet/Hotel/pkg/config"
	"github.com/anonymfrominternet/Hotel/pkg/models"
	"github.com/anonymfrominternet/Hotel/pkg/render"
	"net/http"
)

// Использование Repository pattern
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) MainPage(writer http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Main Page"

	render.RenderTemplate(writer, "main.page.gohtml", &models.TemplateData{StringMap: stringMap})
}
func (m *Repository) AboutPage(writer http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "About Page"
	render.RenderTemplate(writer, "about.page.gohtml", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) PresidentPage(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "president.page.gohtml", &models.TemplateData{})
}
func (m *Repository) BusinessPage(writer http.ResponseWriter, request *http.Request) {

	render.RenderTemplate(writer, "business.page.gohtml", &models.TemplateData{})
}

func (m *Repository) Availability(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "calender.page.gohtml", &models.TemplateData{})
}

func (m *Repository) Contacts(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "contacts.page.gohtml", &models.TemplateData{})
}
func (m *Repository) PersonalData(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "personal-data.page.gohtml", &models.TemplateData{})
}
