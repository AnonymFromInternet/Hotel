package main

import (
	"github.com/anonymfrominternet/Hotel/internal/config"
	"github.com/anonymfrominternet/Hotel/internal/handlers"
	"github.com/anonymfrominternet/Hotel/internal/render"
	"log"
	"net/http"
)

func main() {
	// AppConfig and Repository configuration
	var appConfig config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache in main")
	}

	appConfig.TemplateCache = templateCache
	appConfig.UseCache = false

	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)

	render.NewTemplates(&appConfig)
	// AppConfig and Repository  configuration

	http.HandleFunc("/", handlers.Repo.MainPage)
	http.HandleFunc("/about", handlers.Repo.AboutPage)

	_ = http.ListenAndServe(":3000", nil)
}
