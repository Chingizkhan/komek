package service

import (
	"context"
	"github.com/google/uuid"
	"komek/internal/domain/fundraise/entity"
)

type Repository interface {
	GetByAccountID(ctx context.Context, accountID uuid.UUID) ([]entity.Fundraise, error)
	GetByID(ctx context.Context, id uuid.UUID) (entity.Fundraise, error)
	Create(ctx context.Context, in entity.CreateIn) (entity.Fundraise, error)
	CreateType(ctx context.Context, name string) error
	ListActive(ctx context.Context) ([]entity.Fundraise, error)
	Donate(ctx context.Context, id uuid.UUID, amount int64) error
}
