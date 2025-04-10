package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"komek/internal/domain/account/entity"
	client "komek/internal/domain/client/entity"
	fundraise "komek/internal/domain/fundraise/entity"
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
		GetByOwnerID(ctx context.Context, ownerID uuid.UUID) (entity.Account, error)

		//CreateAccount(ctx context.Context, in dto.CreateAccountIn) (domain.Account, error)
		//InfoAccount(ctx context.Context, accountID uuid.UUID) (out domain.Account, err error)
		//ListAccounts(ctx context.Context, accountIDs []string) (out []domain.Account, err error)
	}

	FundraiseService interface {
		GetByID(ctx context.Context, id uuid.UUID) (fundraise.Fundraise, error)
		GetByAccountID(ctx context.Context, id uuid.UUID) ([]fundraise.Fundraise, error)
		Create(ctx context.Context, in fundraise.CreateIn) (fundraise.Fundraise, error)
		CreateType(ctx context.Context, name string) error
		ListActive(ctx context.Context) ([]fundraise.Fundraise, error)
	}

	ClientService interface {
		GetByID(ctx context.Context, id uuid.UUID) (client.Client, error)
		List(ctx context.Context) (client.Clients, error)
		Create(ctx context.Context, in client.CreateIn) (client client.Client, err error)
	}
)
