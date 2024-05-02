package dto

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"komek/internal/domain"
	"net/http"
)

type (
	// swagger:model
	TransferIn struct {
		// required: true
		FromAccountID uuid.UUID `json:"from_account_id"`
		// required: true
		ToAccountID uuid.UUID `json:"to_account_id"`
		// required: true
		Amount int64 `json:"amount"`
	}

	// swagger:model
	TransferOut struct {
		Transaction domain.Transaction `json:"transaction"`
		FromAccount domain.Account     `json:"from_account"`
		ToAccount   domain.Account     `json:"to_account"`
	}

	// CreateAccountIn defines the request for create_account method
	// swagger:model
	CreateAccountIn struct {
		// required: true
		Owner       uuid.UUID `json:"owner"`
		Balance     int64     `json:"balance"`
		HoldBalance int64     `json:"hold_balance"`
		// required: true
		Country string `json:"country"`
		// required: true
		Currency string `json:"currency"`
	}

	GetAccountIn struct {
		ID uuid.UUID
	}

	// ListAccountsIn - Request to list all accounts connected with user
	// swagger:model
	ListAccountsIn struct {
		// required: true
		// id of user which is gotten via access_token
		UserID uuid.UUID `json:"user_id"`
	}

	// ListAccountsOut - Response to lists all accounts connected with user
	// swagger:model
	ListAccountsOut struct {
		Accounts []domain.Account `json:"accounts"`
	}
)

func (in *CreateAccountIn) ParseAndValidate(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(in)
	if err != nil {
		return fmt.Errorf("decode response body: %w", err)
	}
	defer r.Body.Close()
	return nil
}

func (in *TransferIn) ParseAndValidate(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(in)
	if err != nil {
		return fmt.Errorf("decode response body: %w", err)
	}
	defer r.Body.Close()
	return nil
}

func (in *GetAccountIn) ParseAndValidate(r *http.Request) error {
	idString := chi.URLParam(r, "id")

	id, err := uuid.Parse(idString)
	if err != nil {
		return fmt.Errorf("conv acc id: %w", err)
	}

	in.ID = id
	return nil
}
