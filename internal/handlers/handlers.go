package handlers

import (
	"github.com/anonymfrominternet/Hotel/internal/render"
	"net/http"
)

func MainPage(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, "main.page.tmpl")
}
func AboutPage(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, "about.page.tmpl")
}
