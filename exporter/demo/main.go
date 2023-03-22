package main

import (
	"github.com/menah3m/exporter/demo/pkg/collector"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func init() {
	prometheus.MustRegister(collector.NewNodeCollector())
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("Error occur when start server %v", err)
	}
}
