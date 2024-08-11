package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	// Define a global map to hold metrics
	serviceMetrics = map[string]map[string]prometheus.Gauge{}
)

func initPrometheusMetrics(categories map[string]map[string]string) {
	for category, services := range categories {
		if _, exists := serviceMetrics[category]; !exists {
			serviceMetrics[category] = make(map[string]prometheus.Gauge)
		}
		for serviceName := range services {
			// Create a metric if it doesn't exist already
			if _, exists := serviceMetrics[category][serviceName]; !exists {
				serviceMetrics[category][serviceName] = prometheus.NewGauge(prometheus.GaugeOpts{
					Name:        "arvan_services_status",
					Help:        "arvan services status.",
					ConstLabels: prometheus.Labels{"category": category, "service": serviceName},
				})
				prometheus.MustRegister(serviceMetrics[category][serviceName])
			}
		}
	}
}

func fetchStatus() {
	url := "https://www.arvancloudstatus.ir/"
	res, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching the URL: %v", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Printf("Error: status code %d", res.StatusCode)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Printf("Error reading the document: %v", err)
		return
	}

	// Iterate over each category and its services
	categories := map[string]map[string]string{}
	doc.Find(".services-status>.service-status").Each(func(i int, s *goquery.Selection) {
		firstElement := true
		services := map[string]string{}

		var categoryName string
		var categoryStatus string
		// Log each service under the category with its status
		s.Find(" .service-status--info").Each(func(j int, dd *goquery.Selection) {
			serviceName := strings.TrimSpace(dd.Find(".service-status--name").Text())
			status := strings.TrimSpace(dd.Find(".service-status--status").Text())
			if serviceName != "" && status != "" {
				if firstElement {
					firstElement = false
					categoryName = serviceName
					categoryStatus = status
					return
				}
				categories[categoryName] = services
				services[serviceName] = status

				log.Printf("Service: %s, Status: %s", serviceName, status)
			}
		})
		if len(services) == 0 {
			services[categoryName] = categoryStatus
		}
	})
	initPrometheusMetrics(categories)
	updatePrometheusMetrics(categories)
	//log.Println(categories)
}

func updatePrometheusMetrics(categories map[string]map[string]string) {
	for category, services := range categories {
		for serviceName, status := range services {
			if gauge, exists := serviceMetrics[category][serviceName]; exists {
				if status == "Operational" {
					gauge.Set(1)
				} else if status == "Error" {
					gauge.Set(0)
				}
			}
		}
	}
}

func main() {
	// Start HTTP server for Prometheus to scrape
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Fatal(http.ListenAndServe(":8000", nil))
	}()

	// Periodically fetch status
	for {
		log.Println("Starting fetchStatus function...")
		fetchStatus()
		log.Println("Completed fetchStatus function.")
		time.Sleep(60 * time.Second)
	}
}
