package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"komek/internal/domain/fundraise/entity"
)

type Service struct {
	r Repository
	c Cache
}

func New(r Repository, c Cache) *Service {
	return &Service{r, c}
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (fundraise entity.Fundraise, err error) {
	fundraise, err = s.r.GetByID(ctx, id)
	if err != nil {
		return entity.Fundraise{}, fmt.Errorf("error getting fundraise: %w", err)
	}
	quantity, err := s.c.GetDonorsQuantityByFundraiseID(ctx, id)
	if err != nil {
		return entity.Fundraise{}, err
	}

	fundraise.SupportersQuantity = int64(quantity)

	return fundraise, nil
}

func (s *Service) GetByAccountID(ctx context.Context, id uuid.UUID) ([]entity.Fundraise, error) {
	return s.r.GetByAccountID(ctx, id)
}

func (s *Service) Create(ctx context.Context, in entity.CreateIn) (entity.Fundraise, error) {
	return s.r.Create(ctx, in)
}

func (s *Service) CreateType(ctx context.Context, name string) error {
	return s.r.CreateType(ctx, name)
}

func (s *Service) ListActive(ctx context.Context) ([]entity.Fundraise, error) {
	return s.r.ListActive(ctx)
}

func (s *Service) Donate(ctx context.Context, id uuid.UUID, amount int64, withCache bool) error {
	if err := s.r.Donate(ctx, id, amount); err != nil {
		return fmt.Errorf("error Donate via repo: %w", err)
	}
	if !withCache {
		return nil
	}

	if _, err := s.c.SetDonorsQuantityByFundraiseID(ctx, id); err != nil {
		return fmt.Errorf("could not set donors quantity by fundraise in cache: %w", err)
	}

	return nil
}

func (s *Service) Close(ctx context.Context, id uuid.UUID) error {
	return s.r.SetStatus(ctx, id, false)
}

func (s *Service) IsGoalAchieved(ctx context.Context, id uuid.UUID) (bool, error) {
	fundraise, err := s.GetByID(ctx, id)
	if err != nil {
		return false, fmt.Errorf("error getting fundraise via repo: %w", err)
	}
	if fundraise.Collected < fundraise.Goal {
		return false, nil
	}
	return true, nil
}
