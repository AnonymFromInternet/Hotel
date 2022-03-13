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
	"path/filepath"
	"time"
)

var session *scs.SessionManager
var app config.AppConfig
var pathToTemplates = "../../templates"
var functions = template.FuncMap{}

func getRoutes() http.Handler {

	gob.Register(models.PersonalData{})

	app.InProduction = false

	// State Section
	session = scs.New()
	session.Lifetime = 3 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	// State Section

	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("error in main / tc, err := render.CreateTemplateCache()", err)
	}
	app.TemplateCache = tc
	app.UseCache = true

	repo := NewRepo(&app)
	NewHandlers(repo)

	render.NewTemplates(&app)

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	mux.Use(SessionLoad)

	mux.Get("/", Repo.MainPage)
	mux.Get("/about", Repo.AboutPage)
	mux.Get("/president", Repo.PresidentPage)
	mux.Get("/business", Repo.BusinessPage)

	mux.Get("/calender", Repo.Calender)
	mux.Post("/calender", Repo.PostCalender)
	mux.Post("/calender-json", Repo.CalenderJSON)

	mux.Get("/contacts", Repo.Contacts)

	mux.Get("/personal-data", Repo.PersonalData)
	mux.Post("/personal-data", Repo.PostPersonalData)
	mux.Get("/after-personal-data", Repo.AfterPersonalData)

	fileServer := http.FileServer(http.Dir("../../static/"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func CreateTestTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.gohtml", pathToTemplates))
	if err != nil {
		fmt.Println("error in method RenderTemplateTest / filepath.Glob")
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			fmt.Println("error in method RenderTemplateTest / ts, err")
			return cache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.gohtml", pathToTemplates))
		if err != nil {
			fmt.Println("error in method RenderTemplateTest / matches, err")
			return cache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.gohtml", pathToTemplates))
			if err != nil {
				fmt.Println("error in method RenderTemplateTest / ts, err")
				return cache, err
			}
		}
		cache[name] = ts
	}
	return cache, nil
}
