package handlers

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/anonymfrominternet/Hotel/internal/config"
	"github.com/anonymfrominternet/Hotel/internal/models"
	"github.com/anonymfrominternet/Hotel/internal/render"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var session *scs.SessionManager
var appConfig config.AppConfig
var pathToTemplates = "../../templates/"
var functions = template.FuncMap{}

func getRoutes() http.Handler {
	// Creating Loggers
	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	appConfig.InfoLog = infoLogger

	errorLogger := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	appConfig.ErrorLog = errorLogger
	// Creating Loggers

	// Adding custom data types to scs.SessionManager
	gob.Register(models.Reservation{})
	// Adding custom data types to scs.SessionManager

	// State Management configuration
	session = scs.New()
	session.Lifetime = 3 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.IsInProduction
	// State Management configuration

	// AppConfig and Repository configuration
	appConfig.IsInProduction = false
	appConfig.Session = session

	templateCache, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache in main")
	}

	appConfig.TemplateCache = templateCache
	appConfig.UseCache = true

	repo := NewRepo(&appConfig)
	NewHandlers(repo)

	render.NewRenderer(&appConfig)
	// AppConfig and Repository  configuration

	mux := chi.NewRouter()
	// Adding middlewares
	mux.Use(middleware.Recoverer)
	//mux.Use(NoSurfMiddleware)
	mux.Use(SessionLoadMiddleware)
	// Adding middlewares

	// GET Requests Handlers
	mux.Get("/", Repo.MainPage)
	mux.Get("/about", Repo.AboutPage)
	mux.Get("/generals", Repo.Generals)
	mux.Get("/president", Repo.President)
	mux.Get("/search-availability", Repo.Availability)
	mux.Get("/contact", Repo.Contact)
	mux.Get("/reservation", Repo.Reservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)
	// GET Requests Handlers

	// POST Requests Handlers
	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Post("/search-availability-json", Repo.AvailabilityJSON)
	mux.Post("/reservation", Repo.PostReservation)
	// POST Requests Handlers

	// Adding file server
	fileServer := http.FileServer(http.Dir("../../static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	// Adding file server

	return mux
}

func NoSurfMiddleware(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	// Cookie configuration
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   appConfig.IsInProduction,
		SameSite: http.SameSiteLaxMode,
	})
	// Cookie configuration

	return csrfHandler
}

func SessionLoadMiddleware(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// CreateTemplateCache creates map with templates
func CreateTestTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s*.page.tmpl", pathToTemplates))
	if err != nil {
		fmt.Println("error in render package in TemplateTest() in filepath.Glob()")
		return cache, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		tmpl, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println("error in render package in TemplateTest() in templateSet, err :=")
			return cache, err
		}

		layouts, err := filepath.Glob(fmt.Sprintf("%s*.layout.tmpl", pathToTemplates))
		if err != nil {
			fmt.Println("error in render package in TemplateTest() in layouts, err := filepath.Glob()")
			return cache, err
		}

		if len(layouts) > 0 {
			tmpl, err = tmpl.ParseGlob(fmt.Sprintf("%s*.layout.tmpl", pathToTemplates))
			if err != nil {
				fmt.Println("error in render package in TemplateTest() in if len(layouts) > 0 ")
				return cache, err
			}
		}
		cache[name] = tmpl
	}
	return cache, nil
}
