package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/health", healthHandler)
	r.Route("/cpus", func(r chi.Router) {
		r.Get("/", getCpusHandler)
		//		r.Get("/cpus/{id}", getCpuHandler)
	})

	return r
}
