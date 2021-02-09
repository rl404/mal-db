package middleware

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
)

// RealIP is go-chi wrapper.
func RealIP(h http.Handler) http.Handler {
	return middleware.RealIP(h)
}
