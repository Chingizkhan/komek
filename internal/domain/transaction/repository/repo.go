package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"komek/db/sqlc"
	"komek/internal/domain/transaction/entity"
	"komek/internal/service/transactional"
	"komek/pkg/null_value"
	"komek/pkg/postgres"
)

type Repository struct {
	pool *pgxpool.Pool
	q    *sqlc.Queries
}

func New(pg *postgres.Postgres) *Repository {
	return &Repository{pg.Pool, sqlc.New(pg.Pool)}
}

func (r *Repository) GetByAccountID(ctx context.Context, accountID uuid.UUID) ([]entity.Transaction, error) {
	qtx := r.queries(ctx)

	transactions, err := qtx.GetTransactionByAccountID(ctx, null_value.UUID(accountID))
	if err != nil {
		return nil, fmt.Errorf("get transaction by account id: %w", err)
	}

	return r.transactionsToDomain(transactions), nil
}

func (r *Repository) Create(ctx context.Context, transaction entity.Transaction) (entity.Transaction, error) {
	qtx := r.queries(ctx)

	tr, err := qtx.CreateTransaction(ctx, sqlc.CreateTransactionParams{
		FromAccountID: null_value.UUID(transaction.FromAccountID),
		ToAccountID:   null_value.UUID(transaction.ToAccountID),
		Amount:        transaction.Amount,
	})
	if err != nil {
		if err = checkConstraints(err); err != nil {
			return entity.Transaction{}, err
		}
		return entity.Transaction{}, fmt.Errorf("create transaction: %w", err)
	}

	return r.transactionToDomain(tr), nil
}

func (r *Repository) GetTransactionsByAccounts(ctx context.Context, fromAccountID, toAccountID uuid.UUID) ([]entity.Transaction, error) {
	qtx := r.queries(ctx)

	donations, err := qtx.GetTransactionsByAccounts(ctx, sqlc.GetTransactionsByAccountsParams{
		FromAccountID: null_value.UUID(fromAccountID),
		ToAccountID:   null_value.UUID(toAccountID),
	})
	if err != nil {
		return nil, fmt.Errorf("qtx.GetTransactionsByAccounts: %w", err)
	}

	return r.transactionsToDomain(donations), nil
}

func (r *Repository) GetTotalDonationsAmount(ctx context.Context, accountID uuid.UUID) (int64, error) {
	qtx := r.queries(ctx)

	amount, err := qtx.GetDonationsTotalAmountByAccountID(ctx, null_value.UUID(accountID))
	if err != nil {
		return 0, fmt.Errorf("qtx.GetDonationsTotalAmountByAccountID: %w", err)
	}

	return amount, nil
}

func (r *Repository) queries(ctx context.Context) *sqlc.Queries {
	tx, ok := ctx.Value(transactional.TxKey).(pgx.Tx)
	if ok {
		return r.q.WithTx(tx)
	}
	return r.q
}
