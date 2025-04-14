package server

import (
	"net/http"

	"github.com/denisdefreyne/dflag/stores"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewServer(
	addr string,
	r *prometheus.Registry,
	featureFlagsStore *stores.FeatureFlagsStore,
) *http.Server {
	handler := NewHandler(r, featureFlagsStore)

	srv := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	return &srv
}

func NewHandler(
	r *prometheus.Registry,
	featureFlagsStore *stores.FeatureFlagsStore,
) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/metrics", applyMiddleware(promhttp.HandlerFor(r, promhttp.HandlerOpts{}), "/metrics"))
	mux.Handle("GET /healthz", applyMiddleware(http.HandlerFunc(getHealth), "/healthz"))
	mux.Handle("GET /feature-flags/{name}", applyMiddleware(makeFeatureFlagHandlerFunc(featureFlagsStore), "/feature-flags/:name"))

	return mux
}
