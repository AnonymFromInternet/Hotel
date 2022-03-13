package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

type handler struct{}

func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

}
