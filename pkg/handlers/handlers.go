package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/anonymfrominternet/Hotel/pkg/config"
	"github.com/anonymfrominternet/Hotel/pkg/models"
	"github.com/anonymfrominternet/Hotel/pkg/render"
	"net/http"
)

// Repository pattern
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
	render.RenderTemplate(writer, request, "main.page.gohtml", &models.TemplateData{})
}
func (m *Repository) AboutPage(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "about.page.gohtml", &models.TemplateData{})
}

func (m *Repository) PresidentPage(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "president.page.gohtml", &models.TemplateData{})
}
func (m *Repository) BusinessPage(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "business.page.gohtml", &models.TemplateData{})
}

func (m *Repository) Calender(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "calender.page.gohtml", &models.TemplateData{})
}

// PostCalender - этот метод скорее всего удалить
func (m *Repository) PostCalender(writer http.ResponseWriter, request *http.Request) {
	start := request.Form.Get("start")
	end := request.Form.Get("end")
	writer.Write([]byte(fmt.Sprintf("Starting date is %s and ending date is %s", start, end)))
}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) CalenderJSON(writer http.ResponseWriter, request *http.Request) {
	resp := jsonResponse{true, "Available"}

	out, err := json.MarshalIndent(resp, "", "   ")
	if err != nil {
		fmt.Println("error in handlers / CalenderJSON /  out, err := json.MarshalIndent")
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	_, err = writer.Write(out)
	if err != nil {
		fmt.Println("error in handlers / CalenderJSON / _, err = writer.Write(out)")
		return
	}
}

func (m *Repository) Contacts(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "contacts.page.gohtml", &models.TemplateData{})
}
func (m *Repository) PersonalData(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "personal-data.page.gohtml", &models.TemplateData{})
}
