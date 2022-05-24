package render

import (
	"bytes"
	"fmt"
	"github.com/anonymfrominternet/Hotel/internal/config"
	"github.com/anonymfrominternet/Hotel/internal/models"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var appConfig *config.AppConfig
var pathToTemplates = "../../templates/"

// NewRenderer sets the value of the var appConfig
func NewRenderer(appConfigAsParam *config.AppConfig) {
	appConfig = appConfigAsParam
}

var functions = template.FuncMap{}

func AddDefaultData(templateData *models.TemplateData, request *http.Request) *models.TemplateData {
	templateData.Message = appConfig.Session.PopString(request.Context(), "Message")
	templateData.Error = appConfig.Session.PopString(request.Context(), "Error")
	templateData.Warning = appConfig.Session.PopString(request.Context(), "Warning")
	templateData.CSRFToken = nosurf.Token(request)
	return templateData
}

func Template(w http.ResponseWriter, request *http.Request, tmplName string, templateData *models.TemplateData) error {
	// Get the template cache from the app config
	var templateCache map[string]*template.Template
	var err error

	if appConfig.UseCache {
		templateCache = appConfig.TemplateCache
	} else {
		templateCache, err = CreateTemplateCache()
	}

	tmpl, exist := templateCache[tmplName]
	if !exist {
		log.Fatal("error in render package in Template() in tmpl, exist := templateCache[tmplName]. This template does not exist")
	}

	buf := new(bytes.Buffer)

	AddDefaultData(templateData, request)

	err = tmpl.Execute(buf, templateData)
	if err != nil {
		log.Fatal("error in render package in Template() in err = tmpl.Execute(buf, templateData)", err)
		return err
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal("error in render package in Template() in _, err = buf.WriteTo(w)")
		return err
	}
	return nil
}

// CreateTemplateCache creates map with templates
func CreateTemplateCache() (map[string]*template.Template, error) {
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
