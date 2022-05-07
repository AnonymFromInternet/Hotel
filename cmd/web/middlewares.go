package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// NoSurfMiddleware adds CSRF protection to all POST requests
func NoSurfMiddleware(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	// Cookie configuration
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   appConfig.IsInProduction,
		SameSite: http.SameSiteLaxMode,
	})
	// Cookie configuration

	return csrfHandler
}

func SessionLoadMiddleware(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
