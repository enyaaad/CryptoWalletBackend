package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HTTPRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP Requests",
		}, []string{"method", "endpoint", "status"},
	)

	HTTPRequestsDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_requests_duration_secs",
			Help:    "HTTP requests duration in seconds",
			Buckets: []float64{0.1, 0.5, 1, 2, 5, 10},
		}, []string{"method", "endpoint"},
	)
)
