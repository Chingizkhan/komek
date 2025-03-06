package banking_uc

import (
	"context"
	"github.com/jackc/pgx/v5"
	"komek/internal/domain"
	"komek/internal/dto"
	"komek/internal/usecase"
)

type (
	UseCase struct {
		tr      usecase.Transactional
		account usecase.AccountService
		banking usecase.BankingService
	}

	AccountRepo interface {
		Get(ctx context.Context, tx pgx.Tx, id int64) (domain.Account, error)
		Create(ctx context.Context, tx pgx.Tx, in dto.CreateAccountIn) (domain.Account, error)
	}
)

var txKey = struct{}{}

func New(
	tr usecase.Transactional,
	banking usecase.BankingService,
	account usecase.AccountService,
) *UseCase {
	return &UseCase{
		tr:      tr,
		account: account,
		banking: banking,
	}
}
