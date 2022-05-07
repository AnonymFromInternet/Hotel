package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache       bool
	TemplateCache  map[string]*template.Template
	IsInProduction bool
	Session        *scs.SessionManager
}
