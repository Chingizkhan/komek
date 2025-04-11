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
