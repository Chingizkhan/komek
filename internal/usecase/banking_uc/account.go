package banking_uc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"komek/internal/domain"
	"komek/internal/dto"
)

func (s *UseCase) CreateAccount(ctx context.Context, in dto.CreateAccountIn) (out domain.Account, err error) {
	out, err = s.banking.CreateAccount(ctx, in)
	if err != nil {
		return out, fmt.Errorf("banking service -> create account: %w", err)
	}
	return
}

func (s *UseCase) GetAccount(ctx context.Context, accID uuid.UUID) (out domain.Account, err error) {
	out, err = s.banking.InfoAccount(ctx, accID)
	if err != nil {
		return out, fmt.Errorf("banking service -> info account: %w", err)
	}
	return
}

func (s *UseCase) ListAccounts(ctx context.Context, userID uuid.UUID) (out []domain.Account, err error) {
	//todo: get accountIDs via userID
	out, err = s.banking.ListAccounts(ctx, []string{})
	if err != nil {
		return out, fmt.Errorf("banking service -> list accounts: %w", err)
	}
	return
}
