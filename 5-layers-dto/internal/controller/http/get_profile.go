package http

import (
	"net/http"

	"github.com/golang-school/evolution/5-layers-dto/internal/dto"

	"github.com/golang-school/evolution/5-layers-dto/pkg/render"

	"github.com/go-chi/chi/v5"
)

func (h *Handlers) GetProfile(w http.ResponseWriter, r *http.Request) {
	input := dto.GetProfileInput{
		ID: chi.URLParam(r, "id"),
	}

	output, err := h.profileService.GetProfile(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	render.JSON(w, output, http.StatusOK)
}
