package banking_uc

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lib/pq"
	"komek/db/sqlc"
	"komek/internal/domain"
	"komek/internal/dto"
	"log"
)

func (s *Service) CreateAccount(ctx context.Context, in dto.CreateAccountIn) (domain.Account, error) {
	var (
		acc sqlc.Account
		err error
	)

	err = s.tx.Exec(ctx, func(q *sqlc.Queries) error {
		acc, err = q.CreateAccount(ctx, sqlc.CreateAccountParams{
			Owner: pgtype.UUID{
				Bytes: in.Owner,
				Valid: true,
			},
			Balance:  in.Balance,
			Currency: in.Currency,
		})
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				log.Println("createAccount usecase - ", pqErr.Code.Name())
				return errors.New(pqErr.Code.Name())
			}
			return fmt.Errorf("q.CreateAccount: %w", err)
		}
		return nil
	})
	if err != nil {
		return domain.Account{}, fmt.Errorf("tx.Exec: %w", err)
	}

	return domain.Account{
		ID:        acc.ID,
		Owner:     acc.Owner.Bytes,
		Balance:   acc.Balance,
		Currency:  acc.Currency,
		CreatedAt: acc.CreatedAt.Time,
	}, nil
}

func (s *Service) GetAccount(ctx context.Context, in dto.GetAccountIn) (domain.Account, error) {
	var (
		acc sqlc.Account
		err error
	)

	err = s.tx.Exec(ctx, func(q *sqlc.Queries) error {
		acc, err = q.GetAccount(ctx, in.ID)
		if err != nil {
			return fmt.Errorf("q.GetAccount: %w", err)
		}
		return nil
	})
	if err != nil {
		return domain.Account{}, fmt.Errorf("tx.Exec: %w", err)
	}

	return domain.Account{
		ID:        acc.ID,
		Owner:     acc.Owner.Bytes,
		Balance:   acc.Balance,
		Currency:  acc.Currency,
		CreatedAt: acc.CreatedAt.Time,
	}, nil
}
