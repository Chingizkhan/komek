package service

import (
	"context"
	"github.com/google/uuid"
	"komek/internal/domain/fundraise/entity"
)

type Service struct {
	r Repository
}

func New(r Repository) *Service {
	return &Service{r}
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (entity.Fundraise, error) {
	return s.r.GetByID(ctx, id)
}

func (s *Service) Create(ctx context.Context, in entity.CreateIn) (entity.Fundraise, error) {
	return s.r.Create(ctx, in)
}

func (s *Service) ListActive(ctx context.Context) ([]entity.Fundraise, error) {
	return s.r.ListActive(ctx)
}
