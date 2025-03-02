package entity

import (
	"github.com/google/uuid"
	country "komek/internal/domain/country/entity"
	currency "komek/internal/domain/currency/entity"
	"komek/internal/errs"
	"time"
)

type (
	Account struct {
		// The id of the account
		// required: true
		ID          uuid.UUID         `json:"id"`
		Owner       uuid.UUID         `json:"owner"`
		Balance     int64             `json:"balance"`
		HoldBalance int64             `json:"hold_balance"`
		Status      AccountStatus     `json:"status"`
		Currency    currency.Currency `json:"currency"`
		Country     country.Country   `json:"country"`
		CreatedAt   time.Time         `json:"created_at"`
		UpdatedAt   time.Time         `json:"updated_at"`
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

func (s AccountStatus) Validate() error {
	for _, st := range AccountStatuses {
		if st != s {
			return errs.AccountStatusInvalid
		}
	}
	return nil
}
