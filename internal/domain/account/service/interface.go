package service

import (
	"context"
	"github.com/google/uuid"
	"komek/internal/domain/account/entity"
)

type AccountRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (entity.Account, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (entity.Account, error)
	Create(ctx context.Context, in entity.CreateIn) (entity.Account, error)
}
