package main

import (
	"github.com/anonymfrominternet/Hotel/internal/helpers"
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

// AuthMiddleware checks for a token to check if a user is logged in or not
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !helpers.IsAuthenticated(request) {
			session.Put(request.Context(), "error", "You must be logged in")
			http.Redirect(writer, request, "/user/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(writer, request)
	})
}
