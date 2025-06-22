package api

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

// func getCpuHandler(w http.ResponseWriter, r *http.Request) {
// 	// Call service layer, return JSON
// }

func getCpusHandler(w http.ResponseWriter, r *http.Request) {
	// Call service layer, return JSON
}
