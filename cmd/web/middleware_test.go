package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var handler handler

	h := NoSurf(&handler)

	switch v := h.(type) {
	case http.Handler:
		// nothing
	default:
		t.Error(fmt.Sprintf("TestNoSurf / Type is not http.Handler, but %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var handler handler

	h := NoSurf(&handler)

	switch v := h.(type) {
	case http.Handler:
		// nothing
	default:
		t.Error(fmt.Sprintf("TestSessionLoad / Type is not http.Handler, but %T", v))
	}
}
