package api

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_count",
		Help: "No of request handled by Ping Handler",
	})

func pingHandler(w http.ResponseWriter, r *http.Request) {
	pingCounter.Inc()
	_, err := w.Write([]byte("pong"))
	if err != nil {
		panic(err)
	}
}

func RunServer(Port string) {
	prometheus.MustRegister(pingCounter)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/ping", pingHandler)
	err := http.ListenAndServe(":"+Port, nil)
	if err != nil {
		panic(err)
	}
}
