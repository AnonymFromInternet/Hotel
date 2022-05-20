package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type PostData struct {
	key   string
	value string
}

var testsConfigs = []struct {
	name               string
	url                string
	method             string
	params             []PostData
	expectedStatusCode int
}{
	{name: "main", url: "/", method: "GET", params: []PostData{}, expectedStatusCode: http.StatusOK},
	{name: "about", url: "/about", method: "GET", params: []PostData{}, expectedStatusCode: http.StatusOK},
	{name: "generals", url: "/generals", method: "GET", params: []PostData{}, expectedStatusCode: http.StatusOK},
	{name: "president", url: "/president", method: "GET", params: []PostData{}, expectedStatusCode: http.StatusOK},
	{name: "search-availability", url: "/search-availability", method: "GET", params: []PostData{}, expectedStatusCode: http.StatusOK},
	{name: "contact", url: "/contact", method: "GET", params: []PostData{}, expectedStatusCode: http.StatusOK},
	{name: "reservation", url: "/reservation", method: "GET", params: []PostData{}, expectedStatusCode: http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()

	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close()

	for _, testConfig := range testsConfigs {
		if testConfig.method == "GET" {
			// Requests with GET Method
			response, err := testServer.Client().Get(testServer.URL + testConfig.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if testConfig.expectedStatusCode != response.StatusCode {
				t.Errorf("in %s expected status code is %d, but actual is %d", testConfig.name, testConfig.expectedStatusCode, response.StatusCode)
			}
		} else {
			// Requests with POST Method
		}
	}
}
