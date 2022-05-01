package main

import (
	"fmt"
	"net/http"
)

func MainPage(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "This is main page")
}
func AboutPage(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "This is about page")
}

func main() {
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/about", AboutPage)
	_ = http.ListenAndServe(":3000", nil)
}
