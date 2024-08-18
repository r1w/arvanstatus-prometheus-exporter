package server

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// StartServer initializes the HTTP server for Prometheus metrics and starts the periodic status fetch
func StartServer(fetchStatus func()) {
	// Start HTTP server for Prometheus to scrape
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Fatal(http.ListenAndServe("0.0.0.0:8002", nil))
	}()

	// Periodically fetch status
	for {
		log.Println("Starting fetchStatus function...")
		fetchStatus()
		log.Println("Completed fetchStatus function.")
		time.Sleep(60 * time.Second)
	}
}
