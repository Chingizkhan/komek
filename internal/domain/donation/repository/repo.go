package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"komek/db/sqlc"
	"komek/internal/domain/donation/entity"
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

func (r *Repository) Create(ctx context.Context, in entity.CreateDonationIn) error {
	qtx := r.queries(ctx)

	if err := qtx.CreateDonation(ctx, sqlc.CreateDonationParams{
		FundraiseID:   null_value.UUID(in.FundraiseID),
		TransactionID: null_value.UUID(in.TransactionID),
		ClientID:      null_value.UUID(in.ClientID),
	}); err != nil {
		return fmt.Errorf("qtx.CreateDonation: %w", err)
	}
	return nil
}

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (entity.Donation, error) {
	qtx := r.queries(ctx)

	donation, err := qtx.GetDonationByID(ctx, null_value.UUID(id))
	if err != nil {
		return entity.Donation{}, err
	}

	return r.mapDonation(donation), nil
}

func (r *Repository) GetByTransactionID(ctx context.Context, id uuid.UUID) (entity.Donation, error) {
	qtx := r.queries(ctx)

	donation, err := qtx.GetDonationByTransactionID(ctx, null_value.UUID(id))
	if err != nil {
		return entity.Donation{}, fmt.Errorf("qtx.GetDonationByTransactionID: %w", err)
	}

	return r.mapDonation(sqlc.GetDonationByIDRow(donation)), nil
}

func (r *Repository) queries(ctx context.Context) *sqlc.Queries {
	tx, ok := ctx.Value(transactional.TxKey).(pgx.Tx)
	if ok {
		return r.q.WithTx(tx)
	}
	return r.q
}
