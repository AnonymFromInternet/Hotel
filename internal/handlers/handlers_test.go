package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type postData struct {
	key   string
	value string
}

var tests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"main", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"president", "/president", "GET", []postData{}, http.StatusOK},
	{"business", "/business", "GET", []postData{}, http.StatusOK},
	{"calender", "/calender", "GET", []postData{}, http.StatusOK},
	{"contacts", "/contacts", "GET", []postData{}, http.StatusOK},
	{"pd", "/personal-data", "GET", []postData{}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	testServer := httptest.NewTLSServer(routes)

	defer testServer.Close()

	for _, test := range tests {
		if test.method == "GET" {
			response, err := testServer.Client().Get(testServer.URL + test.url)
			if err != nil {
				fmt.Println("error in handlers_test / TestHandlers / response, err := testServer.Client()")
				t.Error(err)
				t.Fatal(err)
			}
			if response.StatusCode != test.expectedStatusCode {
				t.Errorf("False expected status code")
			}
		} else {

		}
	}
}
