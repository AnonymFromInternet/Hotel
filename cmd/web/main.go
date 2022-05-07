package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/anonymfrominternet/Hotel/internal/config"
	"github.com/anonymfrominternet/Hotel/internal/handlers"
	"github.com/anonymfrominternet/Hotel/internal/render"
	"log"
	"net/http"
	"time"
)

const portNumber = "localhost:3000"

var appConfig config.AppConfig
var session *scs.SessionManager

func main() {
	// AppConfig and Repository configuration
	appConfig.IsInProduction = false
	appConfig.Session = session

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

	// State Management configuration
	session = scs.New()
	session.Lifetime = 3 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.IsInProduction
	// State Management configuration

	// Server configuration
	server := http.Server{
		Addr:    portNumber,
		Handler: routes(&appConfig),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("cannot start listen and serve")
	}
	// Server configuration
}
