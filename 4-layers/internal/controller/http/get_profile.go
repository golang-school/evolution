package http

import (
	"net/http"

	"github.com/golang-school/evolution/4-layers/pkg/render"

	"github.com/go-chi/chi/v5"
)

func (h *Handlers) GetProfile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	output, err := h.profileService.GetProfile(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	render.JSON(w, output, http.StatusOK)
}
