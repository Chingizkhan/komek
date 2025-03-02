package entity

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID        uuid.UUID `json:"id"`
	AccountID uuid.UUID `json:"account_id"`
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
