package word

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	word_db "komek/db/sqlc"
	"komek/internal/domain/word"
	"komek/pkg/postgres"
)

const (
	_repoName         = "word repository"
	_defaultEntityCap = 50
)

type Repository struct {
	pool *pgxpool.Pool
	q    *word_db.Queries
}

func New(pg *postgres.Postgres) *Repository {
	return &Repository{pg.Pool, word_db.New(pg.Pool)}
}

func (r *Repository) Get(ctx context.Context, value string, userID uuid.UUID) (word.Word, error) {
	const fn = "Get"

	w, err := r.q.GetWord(ctx, word_db.GetWordParams{
		Value: value,
		FkUserID: uuid.NullUUID{
			UUID:  userID,
			Valid: true,
		},
	})
	if err != nil {
		return word.Word{}, fmt.Errorf("%s - %s - %w", _repoName, fn, err)
	}
	return word.Word{
		Value:       w.Value,
		Language:    word.Language(w.Language),
		Translation: w.Translation,
	}, nil
}

func (r *Repository) Save(ctx context.Context, word word.Word) error {
	const fn = "Save"

	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return fmt.Errorf("%s - %s - %w", _repoName, fn, err)
	}
	defer tx.Rollback(ctx)

	qtx := r.q.WithTx(tx)

	err = qtx.SaveWord(ctx, word_db.SaveWordParams{
		Value:       word.Value,
		Language:    string(word.Language),
		Translation: word.Translation,
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

	_, err := r.q.UpdateWord(ctx, word_db.UpdateWordParams{
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

	_, err := r.q.DeleteWord(ctx, word_db.DeleteWordParams{
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
