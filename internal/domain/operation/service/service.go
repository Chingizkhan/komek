package service

import (
	"context"
	"github.com/google/uuid"
	"komek/internal/domain/operation/entity"
)

type Service struct {
	r Repository
}

func New(r Repository) *Service {
	return &Service{r: r}
}

func (s *Service) Create(ctx context.Context, in entity.CreateIn) (entity.Operation, error) {
	return s.r.Create(ctx, in)
}

func (s *Service) GetByTransactionID(ctx context.Context, transactionID uuid.UUID) ([]entity.Operation, error) {
	return s.r.GetByTransactionID(ctx, transactionID)
}
