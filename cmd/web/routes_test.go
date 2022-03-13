package main

import (
	"fmt"
	"github.com/anonymfrominternet/Hotel/internal/config"
	"github.com/go-chi/chi/v5"
	"testing"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// nothing
	default:
		t.Error(fmt.Sprintf("TestSessionLoad / Type is not http.Handler, but %T", v))
	}
}
