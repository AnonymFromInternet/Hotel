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

func (repo *Repository) MainPage(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, "main.page.tmpl", &models.TemplateData{})
}
func (repo *Repository) AboutPage(writer http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["defaultData"] = "default value"
	render.Template(writer, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
