package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", MainPage)
	http.HandleFunc("/about", AboutPage)

	_ = http.ListenAndServe(":3000", nil)
}
