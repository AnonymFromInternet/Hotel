package main

import (
	"github.com/anonymfrominternet/Hotel/pkg/config"
	"github.com/anonymfrominternet/Hotel/pkg/handlers"
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

	return mux
}