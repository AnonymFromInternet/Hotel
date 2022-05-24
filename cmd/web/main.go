package main

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/anonymfrominternet/Hotel/internal/config"
	"github.com/anonymfrominternet/Hotel/internal/driver"
	"github.com/anonymfrominternet/Hotel/internal/handlers"
	"github.com/anonymfrominternet/Hotel/internal/helpers"
	"github.com/anonymfrominternet/Hotel/internal/models"
	"github.com/anonymfrominternet/Hotel/internal/render"
	"log"
	"net/http"
	"os"
	"time"
)

const portNumber = "localhost:3000"

var appConfig config.AppConfig
var session *scs.SessionManager

func main() {
	db, err := run()
	if err != nil {
		log.Fatal("error by running the run() function", err)
	}
	defer db.SQL.Close()

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

func run() (*driver.DB, error) {
	// Creating Loggers
	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	appConfig.InfoLog = infoLogger

	errorLogger := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	appConfig.ErrorLog = errorLogger
	// Creating Loggers

	// Adding custom data types to scs.SessionManager
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Reservation{})
	gob.Register(models.RoomRestriction{})
	gob.Register(models.Restriction{})
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

	// Connecting to database
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=hotel user=arturkeil password=")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// Connecting to database

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache in main")
		return nil, err
	}

	appConfig.TemplateCache = templateCache
	appConfig.UseCache = false

	helpers.GetAppConfigToTheHelpersPackage(&appConfig)

	repo := handlers.NewRepo(&appConfig, db)
	handlers.NewHandlers(repo)

	render.NewRenderer(&appConfig)
	// AppConfig and Repository  configuration
	return db, nil
}
