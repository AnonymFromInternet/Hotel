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

	// Requests Handlers
	mux.Get("/", handlers.Repo.MainPage)
	mux.Get("/about", handlers.Repo.AboutPage)
	// Requests Handlers

	return mux
}
