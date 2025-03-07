package entity

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	operation "komek/internal/domain/operation/entity"
	transaction "komek/internal/domain/transaction/entity"
	"net/http"
	"time"
)

type (
	Transaction struct {
		ID            uuid.UUID             `json:"id"`
		FromAccountID uuid.UUID             `json:"from_account_id"`
		ToAccountID   uuid.UUID             `json:"to_account_id"`
		Amount        int64                 `json:"amount"`
		CreatedAt     time.Time             `json:"created_at"`
		Operations    []operation.Operation `json:"operations"`
	}

	TransferIn struct {
		FromUserID    uuid.UUID `json:"from_user_id"`
		ToUserID      uuid.UUID `json:"to_user_id"`
		ToAccountID   uuid.UUID `json:"to_account_id"`
		FromAccountID uuid.UUID `json:"from_account_id"`
		AmountFloat   float64   `json:"amount"`
		Amount        int64
	}

	WithdrawIn struct {
		TransactionID uuid.UUID
		AccountID     uuid.UUID
		Amount        int64
	}

	RefillIn struct {
		TransactionID uuid.UUID
		AccountID     uuid.UUID
		Amount        int64
	}
)

func NewTransaction(tr transaction.Transaction, ops ...operation.Operation) Transaction {
	return Transaction{
		ID:            tr.ID,
		FromAccountID: tr.FromAccountID,
		ToAccountID:   tr.ToAccountID,
		Amount:        tr.Amount,
		CreatedAt:     tr.CreatedAt,
		Operations:    ops,
	}
}

func (req *TransferIn) ParseHttpBody(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return fmt.Errorf("can not decode body: %w", err)
	}
	defer r.Body.Close()
	return nil
}
