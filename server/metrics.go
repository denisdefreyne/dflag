package server

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestsTotalCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Tracks the number of HTTP requests.",
	}, []string{"handler", "method", "code"},
)

var requestDurationHistogram = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Tracks the latencies for HTTP requests.",
		Buckets: prometheus.ExponentialBuckets(0.1, 1.5, 5),
	},
	[]string{"handler", "method", "code"},
)

var requestSizeSummary = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "http_request_size_bytes",
		Help: "Tracks the size of HTTP requests.",
	},
	[]string{"handler", "method", "code"},
)

var responseSizeSummary = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "http_response_size_bytes",
		Help: "Tracks the size of HTTP responses.",
	},
	[]string{"handler", "method", "code"},
)

func applyPrometheusMiddleware(handler http.Handler, name string) http.Handler {
	handler = promhttp.InstrumentHandlerCounter(
		requestsTotalCounter.MustCurryWith(prometheus.Labels{"handler": name}),
		handler,
	)

	handler = promhttp.InstrumentHandlerDuration(
		requestDurationHistogram.MustCurryWith(prometheus.Labels{"handler": name}),
		handler,
	)

	handler = promhttp.InstrumentHandlerRequestSize(
		requestSizeSummary.MustCurryWith(prometheus.Labels{"handler": name}),
		handler,
	)

	handler = promhttp.InstrumentHandlerResponseSize(
		responseSizeSummary.MustCurryWith(prometheus.Labels{"handler": name}),
		handler,
	)

	return handler
}

func RegisterMetrics(r *prometheus.Registry) {
	r.MustRegister(requestsTotalCounter)
	r.MustRegister(requestDurationHistogram)
	r.MustRegister(requestSizeSummary)
	r.MustRegister(responseSizeSummary)
}
