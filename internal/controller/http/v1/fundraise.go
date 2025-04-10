package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"komek/internal/domain/fundraise/entity"
	"komek/pkg/money"
	"net/http"
)

func (h *Handler) fundraiseRoutes(r *chi.Mux) {
	r.Route("/fundraise", func(r chi.Router) {
		// public
		r.Get("/", h.listFundraises)
		r.Get("/{id}", h.getFundraise)
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

type (
	GetOut struct {
		ID          uuid.UUID `json:"id"`
		Name        string    `json:"name"`
		ImageUrl    string    `json:"image_url"`
		City        string    `json:"city"`
		Categories  []string  `json:"categories"`
		Goal        float64   `json:"goal"`
		Collected   float64   `json:"collected"`
		Description string    `json:"description"`
	}
)

func (r *GetOut) FromDomain(fund entity.GetOut) GetOut {
	*r = GetOut{
		ID:          fund.ID,
		Name:        fund.Name,
		ImageUrl:    fund.ImageUrl,
		City:        fund.City,
		Categories:  fund.Categories,
		Goal:        money.ToFloat(fund.Goal),
		Collected:   money.ToFloat(fund.Collected),
		Description: fund.Description,
	}
	return *r
}

func (h *Handler) getFundraise(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	fundraiseID, err := uuid.Parse(idParam)
	if err != nil {
		h.Error(w, err, http.StatusBadRequest, "getFundraise - uuid.Parse")
		return
	}

	fund, err := h.funds.GetByID(r.Context(), fundraiseID)
	if err != nil {
		h.Error(w, err, http.StatusInternalServerError, "getFundraise - h.funds.GetByID")
		return
	}

	h.Resp(w, new(GetOut).FromDomain(fund), http.StatusOK)
}
