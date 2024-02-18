package account_repo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	account_db "komek/db/sqlc"
	"komek/internal/domain"
	"komek/pkg/postgres"
)

type Repository struct {
	pool *pgxpool.Pool
	q    *account_db.Queries
}

func New(pg *postgres.Postgres) *Repository {
	return &Repository{pg.Pool, account_db.New(pg.Pool)}
}

func (r *Repository) Get(ctx context.Context, id int64) (domain.Account, error) {
	acc, err := r.q.GetAccount(ctx, id)
	if err != nil {
		return domain.Account{}, fmt.Errorf("r.q.GetAccount: %w", err)
	}
	return domain.Account{
		ID:        acc.ID,
		Owner:     acc.Owner,
		Balance:   acc.Balance,
		Currency:  acc.Currency,
		CreatedAt: acc.CreatedAt,
	}, nil
}
