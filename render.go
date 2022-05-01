package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmplName string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmplName)
	if err != nil {
		fmt.Println("error in template.ParseFiles()")
		return
	}

	err = parsedTemplate.Execute(w, parsedTemplate)
	if err != nil {
		fmt.Println("error in parsedTemplate.Execute()")
		return
	}
}
