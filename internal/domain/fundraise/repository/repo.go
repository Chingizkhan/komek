package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"komek/db/sqlc"
	"komek/internal/domain/fundraise/entity"
	"komek/internal/errs"
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

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (entity.Fundraise, error) {
	qtx := r.queries(ctx)

	fundraise, err := qtx.GetFundraiseByID(ctx, null_value.UUID(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Fundraise{}, errs.FundraiseNotFound
		}
		return entity.Fundraise{}, fmt.Errorf("r.q.GetFundraiseByID: %w", err)
	}
	return r.mapFundraise(fundraise), nil
}

func (r *Repository) Create(ctx context.Context, in entity.CreateIn) (entity.Fundraise, error) {
	qtx := r.queries(ctx)

	fundraise, err := qtx.CreateFundraise(ctx, sqlc.CreateFundraiseParams{
		Goal:      null_value.Int64(&in.Goal),
		Collected: null_value.Int64(&in.Collected),
		AccountID: null_value.UUID(in.AccountID),
		IsActive:  null_value.Bool(in.IsActive),
	})
	if err != nil {
		return entity.Fundraise{}, fmt.Errorf("r.q.CreateFundraise: %w", err)
	}
	return r.mapFundraise(fundraise), nil
}

func (r *Repository) ListActive(ctx context.Context) ([]entity.Fundraise, error) {
	qtx := r.queries(ctx)

	fundraises, err := qtx.ListActiveFundraises(ctx)
	if err != nil {
		return nil, fmt.Errorf("r.q.ListActiveFundraises: %w", err)
	}
	return r.mapFundraises(fundraises), nil
}

func (r *Repository) queries(ctx context.Context) *sqlc.Queries {
	tx, ok := ctx.Value(transactional.TxKey).(pgx.Tx)
	if ok {
		return r.q.WithTx(tx)
	}
	return r.q
}
