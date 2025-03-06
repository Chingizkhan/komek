package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"komek/internal/domain/account/entity"
	operation "komek/internal/domain/operation/entity"
	banking "komek/internal/service/banking/entity"
)

type (
	Transactional interface {
		Exec(ctx context.Context, fn func(tx pgx.Tx) error) error
		ExecContext(ctx context.Context, fn func(ctx context.Context) error) error
	}

	BankingService interface {
		Transfer(ctx context.Context, in banking.TransferIn) (tr banking.Transaction, err error)
		Withdraw(ctx context.Context, in banking.WithdrawIn) (op operation.Operation, err error)
		Refill(ctx context.Context, in banking.RefillIn) (op operation.Operation, err error)
	}

	AccountService interface {
		GetByID(ctx context.Context, id uuid.UUID) (entity.Account, error)
		Create(ctx context.Context, in entity.CreateIn) (entity.Account, error)
		GetByUserID(ctx context.Context, userID uuid.UUID) (entity.Account, error)

		//CreateAccount(ctx context.Context, in dto.CreateAccountIn) (domain.Account, error)
		//InfoAccount(ctx context.Context, accountID uuid.UUID) (out domain.Account, err error)
		//ListAccounts(ctx context.Context, accountIDs []string) (out []domain.Account, err error)
	}
)
