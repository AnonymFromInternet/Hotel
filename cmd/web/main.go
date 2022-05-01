package main

import (
	"github.com/anonymfrominternet/Hotel/internal/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.MainPage)
	http.HandleFunc("/about", handlers.AboutPage)

	_ = http.ListenAndServe(":3000", nil)
}
