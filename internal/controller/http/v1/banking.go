package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	customMiddleware "komek/internal/controller/http/middleware"
	"komek/internal/domain"
	"komek/internal/dto"
	"net/http"
)

func (h *Handler) bankingRoutes(r *chi.Mux) {
	r.Route("/", func(r chi.Router) {
		// protected
		r.Use(customMiddleware.Auth(h.tokenMaker))

		r.Post("/account/create", h.accountCreate)
		r.Get("/account/{:id}", h.accountGet)
		r.Post("/operation/transfer", h.operationTransfer)
	})
}

// Account returned in the response
// swagger:response accountResponse
type accountResponseWrapper struct {
	// All: account in the system
	// in: body
	Body domain.Account
}

// swagger:parameters AccountCreateRequest
type accountCreateRequestWrapper struct {
	// in: body
	Body dto.CreateAccountIn
}

// swagger:response
type accountCreateResponseWrapper struct {
	// in: body
}

// swagger:route POST /account/create Account AccountCreateRequest
// Creates and returns account connected with User
// responses:
// 200: accountResponse

// accountCreate - creates and returns account connected with User
func (h *Handler) accountCreate(w http.ResponseWriter, r *http.Request) {
	req := dto.CreateAccountIn{}
	if err := req.ParseAndValidate(r); err != nil {
		h.Error(w, err, http.StatusBadRequest, "accountCreate - ParseAndValidate")
		return
	}

	payload := h.payload(r)
	req.Owner = payload.UserID
	account, err := h.banking.CreateAccount(r.Context(), req)
	if err != nil {
		h.Error(w, err, http.StatusInternalServerError, "accountCreate - banking.CreateAccount")
		return
	}

	h.Resp(w, account, http.StatusOK)
}

// swagger:parameters GetInfo
type accountIDWrapper struct {
	// The ID of account
	// in: path
	// required: true
	ID uuid.UUID `json:"id"`
}

// swagger:route GET /account/{id} Account GetInfo
// Returns account info connected with User
// responses:
// 200: accountResponse

// accountGet - Returns account info connected with User
func (h *Handler) accountGet(w http.ResponseWriter, r *http.Request) {
	req := dto.GetAccountIn{}
	if err := req.ParseAndValidate(r); err != nil {
		h.Error(w, err, http.StatusBadRequest, "accountGet - ParseAndValidate")
		return
	}

	account, err := h.banking.GetAccount(r.Context(), req.ID)
	if err != nil {
		h.Error(w, err, http.StatusInternalServerError, "accountGet - banking_uc.CreateAccount")
		return
	}

	h.Resp(w, account, http.StatusOK)
}

// accountGet - Returns account info connected with User
func (h *Handler) accountsList(w http.ResponseWriter, r *http.Request) {
	var req dto.ListAccountsIn
	payload := h.payload(r)
	req.UserID = payload.UserID
	account, err := h.banking.ListAccounts(r.Context(), req.UserID)
	if err != nil {
		h.Error(w, err, http.StatusInternalServerError, "accountGet - banking_uc.CreateAccount")
		return
	}

	h.Resp(w, account, http.StatusOK)
}

// swagger:parameters OperationTransferRequest
type transferRequest struct {
	// in: body
	Body dto.TransferIn
}

// swagger:response OperationTransferResponse
type transferResponse struct {
	// in: body
	Body dto.TransferOut
}

// swagger:route POST /operation/transfer Operations OperationTransferRequest
// Process transfer between two accounts
// responses:
// 200: OperationTransferResponse

// operationTransfer - Process transfer between two accounts
func (h *Handler) operationTransfer(w http.ResponseWriter, r *http.Request) {
	req := dto.TransferIn{}
	if err := req.ParseAndValidate(r); err != nil {
		h.Error(w, err, http.StatusBadRequest, "operationTransfer - ParseAndValidate")
		return
	}

	transfer, err := h.banking.Transfer(r.Context(), req)
	if err != nil {
		h.Error(w, err, http.StatusInternalServerError, "operationTransfer - banking_uc.Transfer")
		return
	}

	h.Resp(w, transfer, http.StatusOK)
}
