package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	customMiddleware "komek/internal/controller/http/middleware"
	accEntity "komek/internal/domain/account/entity"
	donation "komek/internal/domain/donation/entity"
	operation "komek/internal/domain/operation/entity"
	transaction "komek/internal/domain/transaction/entity"
	usrEntity "komek/internal/domain/user/entity"
	"komek/internal/service/banking/entity"
	banking "komek/internal/service/banking/entity"
	"komek/internal/usecase/banking_uc"
	"komek/pkg/money"
	"net/http"
	"time"
)

func (h *Handler) bankingRoutes(r *chi.Mux) {
	r.Route("/", func(r chi.Router) {
		// protected
		r.Use(customMiddleware.Auth(h.tokenMaker))

		r.Post("/account/create", h.accountCreate)
		r.Get("/account/{:id}", h.accountGet)
		r.Get("/account/donations", h.accountGetDonations)
		r.Get("/account/{:id}/donations/total", h.accountGetTotalMoneyDonation)
		r.Post("/operation/transfer", h.operationTransfer)
		r.Post("/operation/donate", h.operationDonate)
	})
}

func (h *Handler) operationDonate(w http.ResponseWriter, r *http.Request) {
	var (
		in  entity.DonateIn
		err error
	)

	if err = in.ParseHttpBody(r); err != nil {
		h.Error(w, err, http.StatusBadRequest, "donateFundraise - ParseHttpBody")
		return
	}

	in.Amount = money.ToInt(in.AmountFloat)

	payload := h.payload(r)

	resp, err := h.user.Get(r.Context(), usrEntity.GetIn{ID: payload.UserID})
	if err != nil {
		h.Error(w, err, http.StatusBadRequest, "donateFundraise - user.Get")
		return
	}
	in.FromAccountID = resp.Account.ID

	if err = h.banking.Donate(r.Context(), in); err != nil {
		h.Error(w, err, http.StatusInternalServerError, "donateFundraise - h.banking.Donate")
		return
	}

	h.Resp(w, map[string]string{
		"status": "ok",
	}, http.StatusOK)
}

// Account returned in the response
// swagger:response accountResponse
type accountResponseWrapper struct {
	// All: account in the system
	// in: body
	//Body domain.Account
}

// swagger:parameters AccountCreateRequest
type accountCreateRequestWrapper struct {
	// in: body
	//Body dto.CreateAccountIn
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
	req := accEntity.CreateIn{}
	if err := req.ParseHttpBody(r); err != nil {
		h.Error(w, err, http.StatusBadRequest, "accountCreate - Parse")
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
	payload := h.payload(r)

	account, err := h.banking.GetAccount(r.Context(), payload.UserID)
	if err != nil {
		h.Error(w, err, http.StatusInternalServerError, "accountGet - banking_uc.CreateAccount")
		return
	}

	h.Resp(w, account, http.StatusOK)
}

type TransactionResponse struct {
	ID            uuid.UUID `json:"id"`
	FromAccountID uuid.UUID `json:"from_account_id"`
	ToAccountID   uuid.UUID `json:"to_account_id"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}

func (h *Handler) transactionToResponse(transactions []transaction.Transaction) []TransactionResponse {
	res := make([]TransactionResponse, 0, len(transactions))
	for _, trans := range transactions {
		res = append(res, TransactionResponse{
			ID:            trans.ID,
			FromAccountID: trans.FromAccountID,
			ToAccountID:   trans.ToAccountID,
			Amount:        money.ToFloat(trans.Amount),
			CreatedAt:     trans.CreatedAt,
		})
	}
	return res
}

type DonationResponse struct {
	ID            uuid.UUID `json:"id"`
	FromAccountID uuid.UUID `json:"from_account_id"`
	ToAccountID   uuid.UUID `json:"to_account_id"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
	ClientName    string    `json:"client_name"`
	ClientPhoto   string    `json:"client_photo"`
}

func (h *Handler) donationsToResponse(transactions []donation.Donation) []DonationResponse {
	res := make([]DonationResponse, 0, len(transactions))
	for _, trans := range transactions {
		res = append(res, DonationResponse{
			ID:            trans.ID,
			FromAccountID: trans.FromAccountID,
			ToAccountID:   trans.ToAccountID,
			Amount:        money.ToFloat(trans.Amount),
			CreatedAt:     trans.CreatedAt,
			ClientName:    trans.ClientName,
			ClientPhoto:   trans.ClientPhoto,
		})
	}
	return res
}

func (h *Handler) accountGetDonations(w http.ResponseWriter, r *http.Request) {
	payload := h.payload(r)

	account, err := h.banking.GetAccountByUserID(r.Context(), payload.UserID)
	if err != nil {
		h.Error(w, err, http.StatusNotFound, "accountGetTransactions - banking_uc.GetAccountByUserID")
		return
	}

	donations, err := h.banking.FindDonations(r.Context(), banking_uc.FindDonationsIn{
		FromAccountID: account.ID,
	})
	if err != nil {
		h.Error(w, err, http.StatusInternalServerError, "accountGetTransactions - banking_uc.FindDonations")
		return
	}

	h.Resp(w, h.donationsToResponse(donations), http.StatusOK)
}

func (h *Handler) accountGetTotalMoneyDonation(w http.ResponseWriter, r *http.Request) {
	payload := h.payload(r)

	account, err := h.banking.GetAccountByUserID(r.Context(), payload.UserID)
	if err != nil {
		h.Error(w, err, http.StatusNotFound, "accountGetTotalMoneyDonation - banking_uc.GetAccountByUserID")
		return
	}

	amount, err := h.banking.GetTotalMoneyDonation(r.Context(), account.ID)
	if err != nil {
		h.Error(w, err, http.StatusInternalServerError, "accountGetTotalMoneyDonation - banking_uc.GetTotalMoneyDonation")
		return
	}

	h.Resp(w, money.ToFloat(amount), http.StatusOK)
}

// accountGet - Returns account info connected with User
//func (h *Handler) accountsList(w http.ResponseWriter, r *http.Request) {
//	var req dto.ListAccountsIn
//	payload := h.payload(r)
//	req.UserID = payload.UserID
//	account, err := h.banking.GetAccountByUserID(r.Context(), req.UserID)
//	if err != nil {
//		h.Error(w, err, http.StatusInternalServerError, "accountGet - banking_uc.CreateAccount")
//		return
//	}
//
//	h.Resp(w, account, http.StatusOK)
//}

type (
	Transaction struct {
		ID            uuid.UUID   `json:"id"`
		FromAccountID uuid.UUID   `json:"from_account_id"`
		ToAccountID   uuid.UUID   `json:"to_account_id"`
		Amount        float64     `json:"amount"`
		CreatedAt     int64       `json:"created_at"`
		Operations    []Operation `json:"operations"`
	}

	Operation struct {
		ID            uuid.UUID      `json:"id"`
		Type          operation.Type `json:"type"`
		Amount        float64        `json:"amount"`
		BalanceBefore float64        `json:"balance_before"`
		BalanceAfter  float64        `json:"balance_after"`
		CreatedAt     int64          `json:"created_at"`
	}
)

func (tr Transaction) FromDomain(transactionDomain banking.Transaction) Transaction {
	return Transaction{
		ID:            transactionDomain.ID,
		FromAccountID: transactionDomain.FromAccountID,
		ToAccountID:   transactionDomain.ToAccountID,
		Amount:        money.ToFloat(transactionDomain.Amount),
		CreatedAt:     transactionDomain.CreatedAt.Unix(),
		Operations:    convertOperations(transactionDomain.Operations),
	}
}

func convertOperations(ops []operation.Operation) []Operation {
	res := make([]Operation, 0, len(ops))
	for _, op := range ops {
		res = append(res, Operation{}.FromDomain(op))
	}
	return res
}

func (op Operation) FromDomain(operationDomain operation.Operation) Operation {
	return Operation{
		ID:            operationDomain.ID,
		Type:          operationDomain.Type,
		Amount:        money.ToFloat(operationDomain.Amount),
		BalanceBefore: money.ToFloat(operationDomain.BalanceBefore),
		BalanceAfter:  money.ToFloat(operationDomain.BalanceAfter),
		CreatedAt:     operationDomain.CreatedAt.Unix(),
	}
}

// swagger:parameters OperationTransferRequest
type transferRequest struct {
	// in: body
	//Body dto.TransferIn
}

// swagger:response OperationTransferResponse
type transferResponse struct {
	// in: body
	//Body dto.TransferOut
}

// swagger:route POST /operation/transfer Operations OperationTransferRequest
// Process transfer between two accounts
// responses:
// 200: OperationTransferResponse

// operationTransfer - Process transfer between two accounts
func (h *Handler) operationTransfer(w http.ResponseWriter, r *http.Request) {
	req := banking.TransferIn{}
	if err := req.ParseHttpBody(r); err != nil {
		h.Error(w, err, http.StatusBadRequest, "operationTransfer - ParseAndValidate")
		return
	}

	payload := h.payload(r)
	req.FromUserID = payload.UserID

	req.Amount = money.ToInt(req.AmountFloat)

	tr, err := h.banking.Transfer(r.Context(), req)
	if err != nil {
		h.Error(w, err, http.StatusInternalServerError, "operationTransfer - banking_uc.Transfer")
		return
	}

	h.Resp(w, Transaction{}.FromDomain(tr), http.StatusOK)
}
