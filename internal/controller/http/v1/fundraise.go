package v1

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) fundraiseRoutes(r *chi.Mux) {
	r.Route("/fundraise", func(r chi.Router) {
		// public
		r.Get("/", h.listFundraises)
	})
}

func (h *Handler) listFundraises(w http.ResponseWriter, r *http.Request) {
	funds, err := h.funds.List(r.Context())
	if err != nil {
		h.Error(w, err, http.StatusInternalServerError, "listFundraises - h.funds.List")
		return
	}

	h.Resp(w, funds, http.StatusOK)
}
