package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Country string

type Currency string

type (
	// Account defines the structure for API Komek
	// swagger:model
	Account struct {
		// The id of the account
		// required: true
		ID          uuid.UUID     `json:"id"`
		Owner       uuid.UUID     `json:"owner"`
		Balance     int64         `json:"balance"`
		HoldBalance int64         `json:"hold_balance"`
		Status      AccountStatus `json:"status"`
		Currency    Currency      `json:"currency"`
		Country     Country       `json:"country"`
		CreatedAt   time.Time     `json:"created_at"`
		UpdatedAt   time.Time     `json:"updated_at"`
	}

	Entry struct {
		ID        int64
		AccountID int64
		// can be negative or positive
		Amount    int64
		CreatedAt time.Time
	}

	Transfer struct {
		ID            int64
		FromAccountID int64
		ToAccountID   int64
		// must be positive
		Amount    int64
		CreatedAt time.Time
	}
)

type AccountStatus string

const (
	AccountStatusActive  = "active"
	AccountStatusBlocked = "blocked"
)

var AccountStatuses = []AccountStatus{
	AccountStatusActive,
	AccountStatusBlocked,
}

var (
	AccountStatusInvalid = errors.New("invalid_account_status")
)

func (s AccountStatus) Validate() error {
	for _, st := range AccountStatuses {
		if st != s {
			return AccountStatusInvalid
		}
	}
	return nil
}
