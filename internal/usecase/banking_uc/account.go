package banking_uc

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"komek/internal/domain"
	"komek/internal/dto"
)

func (s *UseCase) CreateAccount(ctx context.Context, in dto.CreateAccountIn) (domain.Account, error) {
	var (
		acc domain.Account
		err error
	)

	err = s.tr.Exec(ctx, func(tx pgx.Tx) error {
		acc, err = s.account.Create(ctx, tx, in)
		if err != nil {
			return fmt.Errorf("account.Create: %w", err)
		}
		return nil
	})
	if err != nil {
		return domain.Account{}, fmt.Errorf("tx.Exec: %w", err)
	}

	return acc, nil
}

func (s *UseCase) GetAccount(ctx context.Context, in dto.GetAccountIn) (domain.Account, error) {
	var (
		acc domain.Account
		err error
	)

	err = s.tr.Exec(ctx, func(tx pgx.Tx) error {
		acc, err = s.account.Get(ctx, tx, in.ID)
		if err != nil {
			return fmt.Errorf("account.Get: %w", err)
		}
		return nil
	})
	if err != nil {
		return domain.Account{}, fmt.Errorf("tx.Exec: %w", err)
	}

	return acc, nil
}
