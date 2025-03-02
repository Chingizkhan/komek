package entity

import (
	"github.com/google/uuid"
	"time"
)

type (
	Type string

	Operation struct {
		ID            uuid.UUID `json:"id"`
		TransactionID uuid.UUID `json:"transaction_id"`
		Type          Type      `json:"type"`
		Amount        float64   `json:"amount"`
		BalanceBefore int64     `json:"balance_before"`
		BalanceAfter  int64     `json:"balance_after"`
		CreatedAt     time.Time `json:"created_at"`
	}
)

const (
	TypeRefill     Type = "refill"
	TypeWithdraw   Type = "withdraw"
	TypeHold       Type = "hold"
	TypeClear      Type = "clear"
	TypeCommission Type = "commission"
)
