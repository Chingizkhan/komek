package entity

import (
	"github.com/google/uuid"
	"time"
)

type (
	Donation struct {
		ID            uuid.UUID `json:"id"`
		FundraiseID   uuid.UUID `json:"fundraise_id"`
		TransactionID uuid.UUID `json:"transaction_id"`
		ClientID      uuid.UUID `json:"client_id"`
		FromAccountID uuid.UUID `json:"from_account_id"`
		ToAccountID   uuid.UUID `json:"to_account_id"`
		Amount        int64     `json:"amount"`
		CreatedAt     time.Time `json:"created_at"`
		ClientName    string    `json:"client_name"`
		ClientPhoto   string    `json:"client_photo"`
	}

	CreateDonationIn struct {
		FundraiseID   uuid.UUID `json:"fundraise_id"`
		TransactionID uuid.UUID `json:"transaction_id"`
		ClientID      uuid.UUID `json:"client_id"`
	}
)
