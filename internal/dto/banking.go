package dto

import "komek/internal/domain"

type (
	CreateTransferRequest struct {
		FromAccountID int64 `json:"from_account_id"`
		ToAccountID   int64 `json:"to_account_id"`
		Amount        int64 `json:"amount"`
	}

	TransferParams struct {
		FromAccountID int64
		ToAccountID   int64
		Amount        int64
	}

	TransferResult struct {
		Transfer    domain.Transfer
		FromAccount domain.Account
		ToAccount   domain.Account
		FromEntry   domain.Entry
		ToEntry     domain.Entry
	}
)
