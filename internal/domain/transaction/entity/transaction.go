package entity

import (
	"github.com/google/uuid"
	"time"
)

type (
	Transaction struct {
		ID            uuid.UUID `json:"id"`
		FromAccountID uuid.UUID `json:"from_account_id"`
		ToAccountID   uuid.UUID `json:"to_account_id"`
		Amount        int64     `json:"amount"`
		CreatedAt     time.Time `json:"created_at"`
	}

	Donation struct {
		ID            uuid.UUID `json:"id"`
		FromAccountID uuid.UUID `json:"from_account_id"`
		ToAccountID   uuid.UUID `json:"to_account_id"`
		Amount        int64     `json:"amount"`
		CreatedAt     time.Time `json:"created_at"`
		ClientName    string    `json:"client_name"`
		ClientPhoto   string    `json:"client_photo"`
	}
)
