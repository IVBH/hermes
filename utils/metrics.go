package utils

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// TotalRequests ✅ Tracks total requests to Hermes APIs
	TotalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hermes_total_requests",
			Help: "Total number of requests to Hermes APIs",
		},
		[]string{"endpoint", "method"},
	)

	// MessagesPublished ✅ Tracks messages published
	MessagesPublished = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "hermes_messages_published",
			Help: "Total number of messages published",
		},
	)

	// ActiveSubscribers ✅ Tracks active subscribers
	ActiveSubscribers = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "hermes_active_subscribers",
			Help: "Number of active subscribers",
		},
	)
)

// InitMetrics Register all metrics with Prometheus
func InitMetrics() {
	prometheus.MustRegister(TotalRequests, MessagesPublished, ActiveSubscribers)
}
