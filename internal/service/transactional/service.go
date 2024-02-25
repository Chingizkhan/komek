package transactional

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"komek/pkg/postgres"
)

type Transactional struct {
	pool *pgxpool.Pool
}

func New(pg *postgres.Postgres) *Transactional {
	return &Transactional{pg.Pool}
}

func (t *Transactional) Exec(ctx context.Context, fn func(tx pgx.Tx) error) error {
	tx, err := t.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}
	if err = fn(tx); err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return fmt.Errorf("CommitTx: %w", err)
	}

	return nil
}
