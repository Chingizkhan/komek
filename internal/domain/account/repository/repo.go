package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"komek/db/sqlc"
	"komek/internal/domain/account/entity"
	"komek/internal/errs"
	"komek/internal/mapper"
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

	u, err := qtx.GetUserByID(ctx, null_value.UUID(userID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, errs.ErrUserNotFound
		}
		return entity.User{}, fmt.Errorf("r.q.GetUser: %w", err)
	}
	return mapper.ConvUserToDomain(u), nil
}

func (r *Repository) queries(ctx context.Context) *sqlc.Queries {
	tx, ok := ctx.Value(transactional.TxKey).(pgx.Tx)
	if ok {
		return r.q.WithTx(tx)
	}
	return r.q
}
