package http

import (
	"net/http"

	"github.com/golang-school/evolution/8-vertical-slices/internal/profile/get_profile"

	"github.com/golang-school/evolution/8-vertical-slices/internal/profile/create_profile"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Роутинг
func Router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)

	r.Get("/live", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	r.Get("/ready", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })

	r.Handle("/metrics", promhttp.Handler())

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/profile", create_profile.HTTPv1)
			r.Get("/profile/{id}", get_profile.HTTPv1)
		})
	})

	return r
}
