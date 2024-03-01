package domain

import (
	"github.com/google/uuid"
	"time"
)

type (
	Account struct {
		ID        int64
		Owner     uuid.UUID
		Balance   int64
		Currency  string
		CreatedAt time.Time
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
