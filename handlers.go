package main

import "net/http"

func MainPage(writer http.ResponseWriter, request *http.Request) {
	RenderTemplate(writer, "main.page.tmpl")
}
func AboutPage(writer http.ResponseWriter, request *http.Request) {
	RenderTemplate(writer, "about.page.tmpl")
}
