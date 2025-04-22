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

func (r *Repository) GetDonationsByAccounts(ctx context.Context, fromAccountID, toAccountID uuid.UUID) ([]entity.Donation, error) {
	qtx := r.queries(ctx)

	donations, err := qtx.GetDonationsByAccounts(ctx, sqlc.GetDonationsByAccountsParams{
		FromAccountID: null_value.UUID(fromAccountID),
		ToAccountID:   null_value.UUID(toAccountID),
	})
	if err != nil {
		return nil, fmt.Errorf("qtx.GetDonationsByAccounts: %w", err)
	}

	return r.donationsToDomain(donations), nil
}

func (r *Repository) GetDonationsByAccount(ctx context.Context, accountID uuid.UUID) ([]entity.Donation, error) {
	qtx := r.queries(ctx)

	donations, err := qtx.GetDonationsByAccountID(ctx, null_value.UUID(accountID))
	if err != nil {
		return nil, fmt.Errorf("get donation by account id: %w", err)
	}

	res := make([]sqlc.GetDonationsByAccountsRow, 0, len(donations))
	for _, donation := range donations {
		res = append(res, sqlc.GetDonationsByAccountsRow{
			ID:            donation.ID,
			FromAccountID: donation.FromAccountID,
			ToAccountID:   donation.ToAccountID,
			Amount:        donation.Amount,
			CreatedAt:     donation.CreatedAt,
			ClientName:    donation.ClientName,
			ClientPhoto:   donation.ClientPhoto,
		})
	}

	return r.donationsToDomain(res), nil
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
