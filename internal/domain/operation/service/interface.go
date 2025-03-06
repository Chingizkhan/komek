package service

import (
	"context"
	"github.com/google/uuid"
	"komek/internal/domain/operation/entity"
)

type Repository interface {
	Create(ctx context.Context, in entity.CreateIn) (entity.Operation, error)
	GetByTransactionID(ctx context.Context, transactionID uuid.UUID) ([]entity.Operation, error)
}
