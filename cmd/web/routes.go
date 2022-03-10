package main

import (
	"github.com/anonymfrominternet/Hotel/internal/config"
	"github.com/anonymfrominternet/Hotel/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.MainPage)
	mux.Get("/about", handlers.Repo.AboutPage)
	mux.Get("/president", handlers.Repo.PresidentPage)
	mux.Get("/business", handlers.Repo.BusinessPage)

	mux.Get("/calender", handlers.Repo.Calender)
	mux.Post("/calender", handlers.Repo.PostCalender)
	mux.Post("/calender-json", handlers.Repo.CalenderJSON)

	mux.Get("/contacts", handlers.Repo.Contacts)
	mux.Get("/personal-data", handlers.Repo.PersonalData)

	fileServer := http.FileServer(http.Dir("../../static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
