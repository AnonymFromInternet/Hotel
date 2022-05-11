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

	mux.Get("/search-availability-json", handlers.Repo.AvailabilityJSON)
	// GET Requests Handlers

	// POST Requests Handlers
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	// POST Requests Handlers

	// Adding file server
	fileServer := http.FileServer(http.Dir("../../static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	// Adding file server

	return mux
}
