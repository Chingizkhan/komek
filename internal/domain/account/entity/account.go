package entity

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	country "komek/internal/domain/country/entity"
	currency "komek/internal/domain/currency/entity"
	"komek/internal/errs"
	"net/http"
	"time"
)

type (
	Accounts []Account

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

	CreateIn struct {
		Owner    uuid.UUID         `json:"owner"`
		Balance  int64             `json:"balance"`
		Country  country.Country   `json:"country"`
		Currency currency.Currency `json:"currency"`
	}

	AddBalanceIn struct {
		Amount      int64     `json:"amount"`
		HoldBalance int64     `json:"hold_balance"`
		AccountID   uuid.UUID `json:"account_id"`
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

func (req *CreateIn) ParseHttpBody(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return fmt.Errorf("can not decode body: %w", err)
	}
	defer r.Body.Close()
	return nil
}
