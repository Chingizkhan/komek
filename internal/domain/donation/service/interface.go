package service

import (
	"context"
	"github.com/google/uuid"
	"komek/internal/domain/donation/entity"
)

type Repository interface {
	Create(ctx context.Context, in entity.CreateDonationIn) error
	GetByID(ctx context.Context, id uuid.UUID) (entity.Donation, error)
}
