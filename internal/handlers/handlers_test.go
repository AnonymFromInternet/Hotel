package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
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
	{"calender", "/calender", "POST", []postData{
		{key: "start", value: "2022-03-03"},
		{key: "end", value: "2022-03-06"},
	}, http.StatusOK},
	{"calender-json", "/calender-json", "POST", []postData{
		{key: "start", value: "2022-03-03"},
		{key: "end", value: "2022-03-06"},
	}, http.StatusOK},
	{"personal-data", "/personal-data", "POST", []postData{
		{key: "first_name", value: "name"},
		{key: "last_name", value: "surname"},
		{key: "email", value: "email"},
		{key: "phone", value: "3"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	testServer := httptest.NewTLSServer(routes)

	defer testServer.Close()

	for _, test := range tests {
		// Get
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
			// POST
		} else {
			values := url.Values{}
			for _, param := range test.params {
				values.Add(param.key, param.value)
				response, err := testServer.Client().PostForm(testServer.URL+test.url, values)
				if err != nil {
					fmt.Println("error in handlers_test / TestHandlers / // POST")
					t.Fatal(err)
				}
				if response.StatusCode != test.expectedStatusCode {
					t.Errorf("False expected status code")
				}
			}
		}
	}
}
