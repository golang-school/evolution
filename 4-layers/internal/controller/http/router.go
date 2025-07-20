package http

import (
	"net/http"

	"github.com/golang-school/evolution/4-layers/internal/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Роутинг
func Router(service *usecase.Profile) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	r.Get("/live", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	r.Get("/ready", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })

	r.Handle("/metrics", promhttp.Handler())

	handlers := NewHandlers(service)

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/profile", handlers.CreateProfile)
			r.Get("/profile/{id}", handlers.GetProfile)
		})
	})

	return r
}
