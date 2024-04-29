package domain

import (
	"github.com/google/uuid"
	"time"
)

type Operation struct {
	ID            uuid.UUID
	TransactionID uuid.UUID
	AccountID     uuid.UUID
	Type          OperationType
	// must be positive
	Amount            int64
	BalanceBefore     int64
	BalanceAfter      int64
	HoldBalanceBefore int64
	HoldBalanceAfter  int64
	CreatedAt         time.Time
}

type OperationType string

const (
	OperationTypeWithdraw = "withdraw"
	OperationTypeRefill   = "refill"
)
