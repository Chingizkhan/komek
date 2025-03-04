package repository

import (
	"komek/db/sqlc"
	"komek/internal/domain/operation/entity"
)

func (r *Repository) operationsToDomain(operations []sqlc.Operation) []entity.Operation {
	result := make([]entity.Operation, 0, len(operations))
	for _, op := range operations {
		result = append(result, r.operationToDomain(op))
	}
	return result
}

func (r *Repository) operationToDomain(operation sqlc.Operation) entity.Operation {
	return entity.Operation{
		ID:            operation.ID.Bytes,
		TransactionID: operation.TransactionID.Bytes,
		Type:          entity.Type(operation.Type),
		Amount:        operation.Amount,
		BalanceBefore: operation.BalanceBefore,
		BalanceAfter:  operation.BalanceAfter,
		CreatedAt:     operation.CreatedAt.Time,
	}
}
