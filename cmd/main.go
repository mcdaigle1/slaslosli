package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "handler", "code"},
	)
)

func init() {
}

func fatalWithStack(message string, err error) {
	log.Printf("FATAL: %s - %v\nStack Trace:\n%s", message, err, debug.Stack())
	os.Exit(1)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://name-generator-app-service/"

	resp, err := http.Get(url)
	if err != nil {
		fatalWithStack("Error when retrieveing url: "+url, err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fatalWithStack("Error reading response from url: "+url, err)
	}

	fmt.Fprintln(w, "Hello, World ... my name is ... my name is ... my name is: "+string(body))

	// Increment the metric
	requestsTotal.WithLabelValues(r.Method, "/", fmt.Sprintf("%d", 200)).Inc()
}

func main() {
	// register /metrics handler
	http.HandleFunc("/metrics", promhttp.Handler().ServeHTTP)
	// register / handler
	http.HandleFunc("/", helloHandler)

	port := "0.0.0.0:8080"
	fmt.Printf("Starting server at http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
