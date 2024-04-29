package domain

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID         uuid.UUID       `json:"id"`
	AccountID  uuid.UUID       `json:"account_id"`
	Type       TransactionType `json:"type"`
	Amount     int64           `json:"amount"`
	Operations []Operation     `json:"operations"`
	RefundedBy uuid.UUID       `json:"refunded_by"`
	CreatedAt  time.Time       `json:"created_at"`
}

type TransactionType string

const (
	TransactionTypeTransfer = "transfer"
	TransactionTypeRefill   = "refill"
	TransactionTypeWithdraw = "withdraw"
)
