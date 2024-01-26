package user

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	user_db "komek/db/sqlc"
	"komek/internal/domain/user"
	"komek/internal/domain/word"
	"komek/pkg/postgres"
)

const (
	_repoName         = "user repository"
	_defaultEntityCap = 50
)

type Repository struct {
	pool *pgxpool.Pool
	q    *user_db.Queries
}

func New(pg *postgres.Postgres) *Repository {
	return &Repository{pg.Pool, user_db.New(pg.Pool)}
}

func (r *Repository) Get(ctx context.Context, userID uuid.UUID) (user.User, error) {
	const fn = "Get"

	u, err := r.q.GetUser(ctx, userID)
	if err != nil {
		return user.User{}, fmt.Errorf("%s - %s - %w", _repoName, fn, err)
	}
	return user.User{
		Id:        u.ID,
		Name:      user.Name(u.Name),
		Login:     user.Login(u.Login),
		Email:     user.Email(u.Email),
		Phone:     user.Phone(u.Phone),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (r *Repository) Save(ctx context.Context, u user.User) error {
	const fn = "Save"

	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return fmt.Errorf("%s - %s - %w", _repoName, fn, err)
	}
	defer tx.Rollback(ctx)

	qtx := r.q.WithTx(tx)

	err = qtx.SaveUser(ctx, user_db.SaveUserParams{
		ID:    uuid.New(),
		Name:  string(u.Name),
		Login: string(u.Login),
		Email: string(u.Email),
		Phone: string(u.Phone),
	})
	if err != nil {
		return fmt.Errorf("%s - %s - %w", _repoName, fn, err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("%s - %s - %w", _repoName, fn, err)
	}
	return nil
}

func (r *Repository) Update(ctx context.Context, pkey string, w word.Word) (word.Word, error) {
	const fn = "Update"

	_, err := r.q.UpdateWord(ctx, user_db.UpdateWordParams{
		Value:       pkey,
		Value_2:     w.Value,
		Language:    string(w.Language),
		Translation: w.Translation,
	})
	if err != nil {
		return word.Word{}, fmt.Errorf("%s - %s - %w", _repoName, fn, err)
	}
	return word.Word{}, nil
}

func (r *Repository) Delete(ctx context.Context, value string, userID uuid.UUID) error {
	const fn = "Delete"

	_, err := r.q.DeleteWord(ctx, user_db.DeleteWordParams{
		Value: value,
		FkUserID: uuid.NullUUID{
			UUID:  userID,
			Valid: true,
		},
	})
	if err != nil {
		return fmt.Errorf("%s - %s - %w", _repoName, fn, err)
	}
	return nil
}
