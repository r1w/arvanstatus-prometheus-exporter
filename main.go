package main

import (
	"cloud_server_status/exporter"
	"cloud_server_status/server"
)

// StartServer initializes the HTTP server for Prometheus metrics and starts the periodic status fetch
func main() {
	// Start HTTP server for Prometheus to scrape
	server.StartServer()
	exporter.StartFetchStatusJob(60)
}
