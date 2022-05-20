package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

type testHandler struct{}

// Adapting testHandler datatype to http.handler interface
func (testHandler *testHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {}
