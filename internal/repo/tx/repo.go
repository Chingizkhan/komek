package tx

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"komek/db/sqlc"
	"komek/pkg/postgres"
)

type (
	Tx interface {
		Exec(ctx context.Context, fn func(*sqlc.Queries) error) error
	}

	TxPostgres struct {
		db *pgxpool.Pool
	}
)

func NewTX(pg *postgres.Postgres) *TxPostgres {
	return &TxPostgres{
		db: pg.Pool,
	}
}

func (r *TxPostgres) Exec(ctx context.Context, fn func(*sqlc.Queries) error) error {
	tx, err := r.db.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.ReadCommitted})
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}

	q := sqlc.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}
