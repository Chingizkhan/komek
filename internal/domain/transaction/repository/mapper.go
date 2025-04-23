package repository

import (
	"komek/db/sqlc"
	"komek/internal/domain/transaction/entity"
)

func (r *Repository) transactionsToDomain(transactions []sqlc.Transaction) []entity.Transaction {
	result := make([]entity.Transaction, 0, len(transactions))
	for _, transaction := range transactions {
		result = append(result, r.transactionToDomain(transaction))
	}
	return result
}

func (r *Repository) transactionToDomain(transaction sqlc.Transaction) entity.Transaction {
	return entity.Transaction{
		ID:            transaction.ID.Bytes,
		FromAccountID: transaction.FromAccountID.Bytes,
		ToAccountID:   transaction.ToAccountID.Bytes,
		Amount:        transaction.Amount,
		CreatedAt:     transaction.CreatedAt.Time,
	}
}
