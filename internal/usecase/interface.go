package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"komek/internal/domain/account/entity"
	client "komek/internal/domain/client/entity"
	fundraise "komek/internal/domain/fundraise/entity"
	operation "komek/internal/domain/operation/entity"
	transaction "komek/internal/domain/transaction/entity"
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
		Donate(ctx context.Context, id uuid.UUID, amount int64, withCache bool) error
		Close(ctx context.Context, id uuid.UUID) error
		IsGoalAchieved(ctx context.Context, id uuid.UUID) (bool, error)
	}

	ClientService interface {
		GetByID(ctx context.Context, id uuid.UUID) (client.Client, error)
		List(ctx context.Context) (client.Clients, error)
		Create(ctx context.Context, in client.CreateIn) (client client.Client, err error)
	}

	TransactionService interface {
		GetByAccountID(ctx context.Context, accountID uuid.UUID) ([]transaction.Transaction, error)
		Create(ctx context.Context, transaction transaction.Transaction) (transaction.Transaction, error)
		FindTransactionsByAccounts(ctx context.Context, fromAccountID, toAccountID uuid.UUID) ([]transaction.Transaction, error)
		GetTotalDonationsAmount(ctx context.Context, accountID uuid.UUID) (int64, error)
	}
)
