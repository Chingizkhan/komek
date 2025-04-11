package db

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
		return entity.Fundraise{}, fmt.Errorf("qtx.GetFundraiseByID: %w", err)
	}
	return r.mapFundraise(fundraise), nil
}

func (r *Repository) GetByAccountID(ctx context.Context, accountID uuid.UUID) ([]entity.Fundraise, error) {
	qtx := r.queries(ctx)

	fundraises, err := qtx.GetFundraisesByAccountID(ctx, null_value.UUID(accountID))
	if err != nil {
		return nil, fmt.Errorf("qtx.GetFundraisesByAccountID: %w", err)
	}

	return r.mapFundraises(fundraises), nil
}

func (r *Repository) CreateType(ctx context.Context, name string) error {
	qtx := r.queries(ctx)

	if _, err := qtx.CreateFundraiseType(ctx, name); err != nil {
		return fmt.Errorf("qtx.CreateFundraiseType: %w", err)
	}

	return nil
}

func (r *Repository) Create(ctx context.Context, in entity.CreateIn) (entity.Fundraise, error) {
	qtx := r.queries(ctx)

	fundraise, err := qtx.CreateFundraise(ctx, sqlc.CreateFundraiseParams{
		Goal:      in.Goal,
		Collected: in.Collected,
		Type:      null_value.UUID(in.TypeID),
		AccountID: null_value.UUID(in.AccountID),
		IsActive:  null_value.Bool(in.IsActive),
	})
	if err != nil {
		return entity.Fundraise{}, fmt.Errorf("qtx.CreateFundraise: %w", err)
	}
	return r.mapFundraise(fundraise), nil
}

func (r *Repository) ListActive(ctx context.Context) ([]entity.Fundraise, error) {
	qtx := r.queries(ctx)

	fundraises, err := qtx.ListActiveFundraises(ctx)
	if err != nil {
		return nil, fmt.Errorf("qtx.ListActiveFundraises: %w", err)
	}
	return r.mapFundraises(fundraises), nil
}

func (r *Repository) Donate(ctx context.Context, id uuid.UUID, amount int64) (err error) {
	qtx := r.queries(ctx)

	if err = qtx.DonateFundraise(ctx, sqlc.DonateFundraiseParams{
		Amount: amount,
		ID:     null_value.UUID(id),
	}); err != nil {
		return fmt.Errorf("qtx.DonateFundraise: %w", err)
	}

	return nil
}

func (r *Repository) queries(ctx context.Context) *sqlc.Queries {
	tx, ok := ctx.Value(transactional.TxKey).(pgx.Tx)
	if ok {
		return r.q.WithTx(tx)
	}
	return r.q
}
