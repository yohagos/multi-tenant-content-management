package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HttpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	HttpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	HttpRequestsInFlight = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_requests_in_flight",
			Help: "Current number of http requests being processed",
		},
	)

	TenantsTotal = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "tenants_total",
			Help: "Total number of tenants",
		},
	)

	TenantsActive = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "tenants_active",
			Help: "Total number of active tenants",
		},
	)

	ContentsTotal = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "contents_total",
			Help: "Total number of contents by tenant",
		},
		[]string{"tenant_id", "status"},
	)

	CacheHitsTotal = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "cache_hits_total",
			Help: "Total number of cache hits",
		},
	)

	CacheMissesTotal = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "cache_misses_total",
			Help: "Total number of cache misses",
		},
	)

	DbQueryDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "db_query_duration_seconds",
			Help:    "Database query duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"operation", "table"},
	)

	DbQueryTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name:    "db_query_total",
			Help:    "Total number of database queries",
		},
		[]string{"operation", "table", "status"},
	)

	RabbitMqPublishTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "rabbitmq_publish_total",
			Help: "Total number of RabbitMQ publishes",
		},
		[]string{"exchange", "routing_key", "status"},
	)
)
