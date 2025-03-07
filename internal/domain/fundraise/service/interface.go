package service

import (
	"context"
	"github.com/google/uuid"
	"komek/internal/domain/fundraise/entity"
)

type Repository interface {
	GetByID(ctx context.Context, id uuid.UUID) (entity.Fundraise, error)
	Create(ctx context.Context, in entity.CreateIn) (entity.Fundraise, error)
	ListActive(ctx context.Context) ([]entity.Fundraise, error)
}
