package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Country string

type Currency string

type (
	Account struct {
		ID          uuid.UUID
		Owner       uuid.UUID
		Balance     int64
		HoldBalance int64
		Status      AccountStatus
		Currency    Currency
		Country     Country
		CreatedAt   time.Time
		UpdatedAt   time.Time
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
