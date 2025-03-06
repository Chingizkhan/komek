package dto

import (
	"github.com/google/uuid"
	"komek/internal/domain"
)

type (
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
