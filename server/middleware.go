package server

import (
	"net/http"

	chi_middleware "github.com/go-chi/chi/v5/middleware"
)

func applyMiddleware(handler http.Handler, name string) http.Handler {
	handler = chi_middleware.Recoverer(handler)
	handler = chi_middleware.Logger(handler)
	handler = chi_middleware.RealIP(handler)
	handler = chi_middleware.RequestID(handler)

	handler = applyPrometheusMiddleware(handler, name)

	return handler
}
