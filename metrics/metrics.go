package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    TaskCreationCounter = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "task_creation_total",
            Help: "Total number of tasks created",
        },
        []string{"status", "method", "route"},
    )

    RequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "Histogram of HTTP request duration in seconds",
            Buckets: prometheus.DefBuckets,
        },
        []string{"status", "method", "route"},
    )
)

func init() {
    prometheus.MustRegister(TaskCreationCounter)
    prometheus.MustRegister(RequestDuration)
}


func MetricsHandler() http.Handler {
    return promhttp.Handler()
}
