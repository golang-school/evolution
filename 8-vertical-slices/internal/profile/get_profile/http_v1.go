package get_profile

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/golang-school/evolution/8-vertical-slices/pkg/render"
)

var usecase *Usecase

func HTTPv1(w http.ResponseWriter, r *http.Request) {
	input := Input{
		ID: chi.URLParam(r, "id"),
	}

	output, err := usecase.GetProfile(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	render.JSON(w, output, http.StatusOK)
}
