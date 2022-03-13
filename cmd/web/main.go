package main

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/anonymfrominternet/Hotel/internal/config"
	"github.com/anonymfrominternet/Hotel/internal/handlers"
	"github.com/anonymfrominternet/Hotel/internal/models"
	"github.com/anonymfrominternet/Hotel/internal/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":3000"

var app config.AppConfig
var session *scs.SessionManager

func main() {
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

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error in main / tc, err := render.CreateTemplateCache()", err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("error in method main / srv.ListenAndServe")
	}
}
