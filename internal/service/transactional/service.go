package transactional

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"komek/pkg/postgres"
)

const TxKey = "tx"

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

func (t *Transactional) ExecContext(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := t.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}
	ctx = context.WithValue(ctx, TxKey, tx)

	if err = fn(ctx); err != nil {
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
