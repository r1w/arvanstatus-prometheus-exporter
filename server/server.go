package server

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func StartServer() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Fatal(http.ListenAndServe("0.0.0.0:8002", nil))
	}()
}
