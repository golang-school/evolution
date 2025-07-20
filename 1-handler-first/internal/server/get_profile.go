package server

import (
	"net/http"

	"github.com/golang-school/evolution/1-handler-first/pkg/render"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (h *Handlers) GetProfile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	profileID, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, ErrUUIDInvalid.Error(), http.StatusBadRequest)

		return
	}

	profile, err := h.postgres.GetProfile(r.Context(), profileID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	render.JSON(w, profile, http.StatusOK)
}
