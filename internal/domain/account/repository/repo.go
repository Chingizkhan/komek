package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"komek/db/sqlc"
	"komek/internal/domain/account/entity"
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

func (r *Repository) GetByID(ctx context.Context, id uuid.UUID) (entity.Account, error) {
	qtx := r.queries(ctx)

	acc, err := qtx.GetAccount(ctx, null_value.UUID(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Account{}, errs.AccountNotFound
		}
		return entity.Account{}, fmt.Errorf("r.q.GetAccount: %w", err)
	}
	return r.mapAccount(acc), nil
}

func (r *Repository) GetByUserID(ctx context.Context, userID uuid.UUID) (entity.Account, error) {
	qtx := r.queries(ctx)

	acc, err := qtx.GetAccountsByUserID(ctx, null_value.UUID(userID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Account{}, errs.AccountNotFound
		}
		return entity.Account{}, fmt.Errorf("r.q.GetAccountsByUserID: %w", err)
	}

	return r.mapAccount(acc), nil
}

func (r *Repository) Create(ctx context.Context, in entity.CreateIn) (entity.Account, error) {
	qtx := r.queries(ctx)

	account, err := qtx.CreateAccount(ctx, CreateAccountRequest{in: in}.toSqlc())
	if err != nil {
		return entity.Account{}, fmt.Errorf("r.q.CreateAccount: %w", err)
	}

	return r.mapAccount(account), nil
}

func (r *Repository) queries(ctx context.Context) *sqlc.Queries {
	tx, ok := ctx.Value(transactional.TxKey).(pgx.Tx)
	if ok {
		return r.q.WithTx(tx)
	}
	return r.q
}
