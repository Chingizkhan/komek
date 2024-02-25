package v1

import (
	"github.com/go-chi/chi/v5"
	"komek/internal/dto"
	"komek/pkg/logger"
	"net/http"
)

func (h *Handler) bankingRoutes(r *chi.Mux) {
	r.Route("/", func(r chi.Router) {
		// protected
		// todo: use middleware with jwt auth
		r.Post("/account/create", h.accountCreate)
		r.Get("/account/{:id}", h.accountGet)
		r.Post("/operation/transfer", h.operationTransfer)
	})
}

func (h *Handler) accountCreate(w http.ResponseWriter, r *http.Request) {
	req := dto.CreateAccountIn{}
	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("accountCreate - ParseAndValidate", logger.Err(err))
		h.Err(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := h.banking.CreateAccount(r.Context(), req)
	if err != nil {
		h.l.Error("accountCreate - banking_uc.CreateAccount", logger.Err(err))
		h.Err(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Resp(w, account, http.StatusOK)
}

func (h *Handler) accountGet(w http.ResponseWriter, r *http.Request) {
	req := dto.GetAccountIn{}
	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("accountGet - ParseAndValidate", logger.Err(err))
		h.Err(w, err.Error(), http.StatusBadRequest)
		return
	}

	account, err := h.banking.GetAccount(r.Context(), req)
	if err != nil {
		h.l.Error("accountGet - banking_uc.CreateAccount", logger.Err(err))
		h.Err(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Resp(w, account, http.StatusOK)
}

func (h *Handler) operationTransfer(w http.ResponseWriter, r *http.Request) {
	req := dto.TransferIn{}
	if err := req.ParseAndValidate(r); err != nil {
		h.l.Error("operationTransfer - ParseAndValidate", logger.Err(err))
		h.Err(w, err.Error(), http.StatusBadRequest)
		return
	}

	transfer, err := h.banking.Transfer(r.Context(), req)
	if err != nil {
		h.l.Error("operationTransfer - banking_uc.Transfer", logger.Err(err))
		h.Err(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Resp(w, transfer, http.StatusOK)
}
