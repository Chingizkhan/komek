package domain

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID         uuid.UUID
	AccountID  uuid.UUID
	Type       TransactionType
	Amount     int64
	Operations []Operation
	RefundedBy uuid.UUID
	CreatedAt  time.Time
}

type TransactionType string

const (
	TransactionTypeTransfer = "transfer"
	TransactionTypeRefill   = "refill"
	TransactionTypeWithdraw = "withdraw"
)
