package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurfMiddleware(t *testing.T) {
	var testHandler *testHandler

	handler := NoSurfMiddleware(testHandler)
	switch dt := handler.(type) {
	case http.Handler:
	//
	default:
		t.Error(fmt.Sprintf("data type is not http.handler, but %T", dt))
	}

}

func TestSessionLoadMiddleware(t *testing.T) {
	var testHandler *testHandler

	handler := SessionLoadMiddleware(testHandler)
	switch dt := handler.(type) {
	case http.Handler:
	//
	default:
		t.Error(fmt.Sprintf("data type is not http.handler, but %T", dt))
	}
}
