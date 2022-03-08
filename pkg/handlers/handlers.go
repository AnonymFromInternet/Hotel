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
