package endpoints

import (
	"errors"
	"github.com/go-chi/render"
	"ms-go-notification/internal/contract"
	"ms-go-notification/internal/internal_errors"
	"net/http"
)

func (h Handler) CreateCampaign(w http.ResponseWriter, r *http.Request) {

	var request contract.NewCampaignDTO
	err := render.DecodeJSON(r.Body, &request)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		return
	}

	id, err := h.CampaignService.Create(request)
	if err != nil {
		if errors.Is(err, internal_errors.ErrInternal) {
			render.Status(r, http.StatusInternalServerError)
		} else {
			render.Status(r, http.StatusBadRequest)
		}
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]string{"id": id})

}
