package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func Template(w http.ResponseWriter, tmplName string) {
	templateCache, err := createTemplateCache()
	if err != nil {
		log.Fatal("error in render package in Template() in templateCache, err := createTemplateCache()")
	}

	tmpl, exist := templateCache[tmplName]
	if !exist {
		log.Fatal("error in render package in Template() in tmpl, exist := templateCache[tmplName]. This template does not exist")
	}

	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, nil)
	if err != nil {
		log.Fatal("error in render package in Template() in err = tmpl.Execute(buf, nil)")
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal("error in render package in Template() in _, err = buf.WriteTo(w)")
	}
}

// createTemplateCache creates map with templates
func createTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../templates/*.page.tmpl")
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

		layouts, err := filepath.Glob("../../templates/*.layout.tmpl")
		if err != nil {
			fmt.Println("error in render package in TemplateTest() in layouts, err := filepath.Glob()")
			return cache, err
		}

		if len(layouts) > 0 {
			tmpl, err = tmpl.ParseGlob("../../templates/*.layout.tmpl")
			if err != nil {
				fmt.Println("error in render package in TemplateTest() in if len(layouts) > 0 ")
				return cache, err
			}
		}
		cache[name] = tmpl
	}
	return cache, nil
}
