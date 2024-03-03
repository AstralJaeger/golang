package main

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	myMetric = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "my_metric",
			Help: "My metric",
		})
)

func init() {
	prometheus.MustRegister(myMetric)
}

func main() {

	go func() {
		log.Println("Starting metrics server on :2112")
		http.Handle("/metrics", promhttp.Handler())
		_ = http.ListenAndServe(":2112", nil)
	}()

	println("metrics server on :2112")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Printf("Waiting for ctx: %v", ctx)

	<-ctx.Done()
}
