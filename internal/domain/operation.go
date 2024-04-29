package domain

import (
	"github.com/google/uuid"
	"time"
)

type Operation struct {
	ID            uuid.UUID     `json:"id"`
	TransactionID uuid.UUID     `json:"transaction_id"`
	AccountID     uuid.UUID     `json:"account_id"`
	Type          OperationType `json:"type"`
	// must be positive
	Amount            int64     `json:"amount"`
	BalanceBefore     int64     `json:"balance_before"`
	BalanceAfter      int64     `json:"balance_after"`
	HoldBalanceBefore int64     `json:"hold_balance_before"`
	HoldBalanceAfter  int64     `json:"hold_balance_after"`
	CreatedAt         time.Time `json:"created_at"`
}

type OperationType string

const (
	OperationTypeWithdraw = "withdraw"
	OperationTypeRefill   = "refill"
)
