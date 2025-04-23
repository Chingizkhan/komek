package service

import (
	"context"
	"github.com/google/uuid"
	"komek/internal/domain/transaction/entity"
)

type Repository interface {
	GetByAccountID(ctx context.Context, accountID uuid.UUID) ([]entity.Transaction, error)
	Create(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error)
	GetTransactionsByAccounts(ctx context.Context, fromAccountID, toAccountID uuid.UUID) ([]entity.Transaction, error)
	GetTotalDonationsAmount(ctx context.Context, accountID uuid.UUID) (int64, error)
}

type Service struct {
	r Repository
}

func New(r Repository) *Service {
	return &Service{r: r}
}

func (s *Service) GetByAccountID(ctx context.Context, accountID uuid.UUID) ([]entity.Transaction, error) {
	return s.r.GetByAccountID(ctx, accountID)
}

func (s *Service) Create(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error) {
	return s.r.Create(ctx, transaction)
}

func (s *Service) FindTransactionsByAccounts(ctx context.Context, fromAccountID, toAccountID uuid.UUID) ([]entity.Transaction, error) {
	return s.r.GetTransactionsByAccounts(ctx, fromAccountID, toAccountID)
}

func (s *Service) GetTotalDonationsAmount(ctx context.Context, accountID uuid.UUID) (int64, error) {
	return s.r.GetTotalDonationsAmount(ctx, accountID)
}
