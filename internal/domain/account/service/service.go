package service

import (
	"context"
	"github.com/google/uuid"
	"komek/internal/domain/account/entity"
)

type Service struct {
	r AccountRepository
}

func New(r AccountRepository) *Service {
	return &Service{r: r}
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (entity.Account, error) {
	return s.r.GetByID(ctx, id)
}

func (s *Service) GetByOwnerID(ctx context.Context, ownerID uuid.UUID) (entity.Account, error) {
	return s.r.GetByOwnerID(ctx, ownerID)
}

func (s *Service) Create(ctx context.Context, in entity.CreateIn) (entity.Account, error) {
	return s.r.Create(ctx, in)
}

func (s *Service) AddBalance(ctx context.Context, in entity.AddBalanceIn) (acc entity.Account, err error) {
	return s.r.AddBalance(ctx, in)
}
