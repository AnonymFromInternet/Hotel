package main

import (
	"encoding/gob"
	"fmt"
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

const portNumber = ":3000"
const dsn = "host=localhost port=5432 dbname=hotel user=arturkeil password="

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	connection, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer connection.SQL.Close()

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("error in method main / srv.ListenAndServe")
	}
}
func run() (*driver.DB, error) {
	gob.Register(models.PersonalData{})

	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// State Section
	session = scs.New()
	session.Lifetime = 3 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	// State Section

	// Connect to database
	fmt.Println("Connecting to database...")
	connection, err := driver.ConnectSQL(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database: in main run() driver.ConnectSQL()")
	}

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error in main / tc, err := render.CreateTemplateCache()", err)
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, connection)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	helpers.NewHelpers(&app)

	return connection, nil
}
