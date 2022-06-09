package config

import (
	"github.com/alexedwards/scs/v2"
	"github.com/anonymfrominternet/Hotel/internal/models"
	"html/template"
	"log"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache       bool
	TemplateCache  map[string]*template.Template
	IsInProduction bool
	Session        *scs.SessionManager
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
	MailChan       chan models.MailData
}
