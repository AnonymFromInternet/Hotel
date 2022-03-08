package render

import (
	"bytes"
	"fmt"
	"github.com/anonymfrominternet/Hotel/pkg/config"
	"github.com/anonymfrominternet/Hotel/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("template with this name does not exist")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error in method RenderTemplate / buf.WriteTo")
		return
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../templates/*.page.gohtml")
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

		matches, err := filepath.Glob("../../templates/*.layout.gohtml")
		if err != nil {
			fmt.Println("error in method RenderTemplateTest / matches, err")
			return cache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../templates/*.layout.gohtml")
			if err != nil {
				fmt.Println("error in method RenderTemplateTest / ts, err")
				return cache, err
			}
		}
		cache[name] = ts
	}
	return cache, nil
}
