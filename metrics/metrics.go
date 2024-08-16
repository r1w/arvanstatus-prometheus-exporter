package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Define a global map to hold metrics
var ServiceMetrics = map[string]map[string]prometheus.Gauge{}

// Initialize Prometheus metrics for given categories
func InitPrometheusMetrics(categories map[string]map[string]string) {
	for category, services := range categories {
		if _, exists := ServiceMetrics[category]; !exists {
			ServiceMetrics[category] = make(map[string]prometheus.Gauge)
		}
		for serviceName := range services {
			// Create a metric if it doesn't exist already
			if _, exists := ServiceMetrics[category][serviceName]; !exists {
				ServiceMetrics[category][serviceName] = prometheus.NewGauge(prometheus.GaugeOpts{
					Name:        "arvan_services_status",
					Help:        "arvan services status.",
					ConstLabels: prometheus.Labels{"category": category, "service": serviceName},
				})
				prometheus.MustRegister(ServiceMetrics[category][serviceName])
			}
		}
	}
}

// Update Prometheus metrics based on the fetched status
func UpdatePrometheusMetrics(categories map[string]map[string]string) {
	for category, services := range categories {
		for serviceName, status := range services {
			if gauge, exists := ServiceMetrics[category][serviceName]; exists {
				if status == "Operational" {
					gauge.Set(1)
				} else if status == "Error" {
					gauge.Set(0)
				}
			}
		}
	}
}
