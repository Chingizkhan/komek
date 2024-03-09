package banking_uc

import (
	"context"
	"github.com/jackc/pgx/v5"
	"komek/db/sqlc"
	"komek/internal/domain"
	"komek/internal/dto"
	"komek/internal/usecase"
)

type (
	UseCase struct {
		tr      usecase.Transactional
		tx      Store
		account AccountRepo
	}

	Store interface {
		Exec(ctx context.Context, fn func(*sqlc.Queries) error) error
	}

	AccountRepo interface {
		Get(ctx context.Context, tx pgx.Tx, id int64) (domain.Account, error)
		Create(ctx context.Context, tx pgx.Tx, in dto.CreateAccountIn) (domain.Account, error)
	}
)

var txKey = struct{}{}

func New(tr usecase.Transactional, store Store, account AccountRepo) *UseCase {
	return &UseCase{
		tr,
		store,
		account,
	}
}
