package service

import (
	"context"
	"github.com/google/uuid"
	"komek/internal/domain/donation/entity"
)

type Service struct {
	r Repository
}

func New(r Repository) *Service {
	return &Service{r: r}
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (entity.Donation, error) {
	return s.r.GetByID(ctx, id)
}

func (s *Service) Create(ctx context.Context, in entity.CreateDonationIn) error {
	return s.r.Create(ctx, in)
}
