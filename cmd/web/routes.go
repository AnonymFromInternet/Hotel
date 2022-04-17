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

	mux.Get("/calender", handlers.Repo.Calendar)
	mux.Post("/calender", handlers.Repo.PostCalendar)
	mux.Post("/calender-json", handlers.Repo.CalendarJSON)

	mux.Get("/contacts", handlers.Repo.Contacts)

	mux.Get("/personal-data", handlers.Repo.Reservation)
	mux.Post("/personal-data", handlers.Repo.PostReservation)
	mux.Get("/after-personal-data", handlers.Repo.ReservationSummary)

	fileServer := http.FileServer(http.Dir("../../static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
