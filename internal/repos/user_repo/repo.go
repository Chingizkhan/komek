package user_repo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	user_db "komek/db/sqlc"
	"komek/internal/domain"
	"komek/pkg/postgres"
)

const (
	_defaultEntityCap = 50
)

type Repository struct {
	pool *pgxpool.Pool
	q    *user_db.Queries
}

func New(pg *postgres.Postgres) *Repository {
	return &Repository{pg.Pool, user_db.New(pg.Pool)}
}

func (r *Repository) Get(ctx context.Context, userID uuid.UUID) (domain.User, error) {
	u, err := r.q.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("r.q.GetUser :%w", err)
	}
	return domain.User{
		ID:            u.ID,
		Name:          u.Name.String,
		Phone:         domain.Phone(u.Phone.String),
		Login:         u.Login,
		EmailVerified: u.EmailVerified.Bool,
		PasswordHash:  u.PasswordHash,
		Email:         domain.Email(u.Email.String),
		CreatedAt:     u.CreatedAt.Time,
		UpdatedAt:     u.UpdatedAt.Time,
	}, nil
}

func (r *Repository) Save(ctx context.Context, u domain.User) error {
	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return fmt.Errorf("r.pool.BeginTx: %w", err)
	}
	defer tx.Rollback(ctx)

	qtx := r.q.WithTx(tx)

	err = qtx.SaveUser(ctx, user_db.SaveUserParams{
		Login:        u.Login,
		PasswordHash: u.PasswordHash,
	})
	if err != nil {
		return fmt.Errorf("qtx.SaveUser: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("tx.Commit: %w", err)
	}
	return nil
}
