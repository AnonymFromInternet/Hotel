package main

import (
	"github.com/anonymfrominternet/Hotel/internal/config"
	"github.com/anonymfrominternet/Hotel/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes(appConfig *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// Adding middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurfMiddleware)
	mux.Use(SessionLoadMiddleware)
	// Adding middlewares

	// GET Requests Handlers
	mux.Get("/", handlers.Repo.MainPage)
	mux.Get("/about", handlers.Repo.AboutPage)
	mux.Get("/generals", handlers.Repo.Generals)
	mux.Get("/president", handlers.Repo.President)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/reservation", handlers.Repo.Reservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)
	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoomWithId)
	mux.Get("/user/login", handlers.Repo.Login)
	mux.Get("/user/logout", handlers.Repo.Logout)
	// GET Requests Handlers

	// POST Requests Handlers
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Post("/reservation", handlers.Repo.PostReservation)
	mux.Post("/user/login", handlers.Repo.PostLogin)
	// POST Requests Handlers

	// Secure pages Section
	mux.Route("/admin", func(r chi.Router) {
		r.Use(AuthMiddleware)

		r.Get("/dashboard", handlers.Repo.AdminDashboard)
	})
	// Secure pages Section

	// Adding file server
	fileServer := http.FileServer(http.Dir("../../static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	// Adding file server

	return mux
}
