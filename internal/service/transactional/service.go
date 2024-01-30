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

func (t *Transactional) Start(ctx context.Context) (pgx.Tx, error) {
	tx, err := t.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return nil, fmt.Errorf("t.pool.BeginTx: %w", err)
	}
	return tx, nil
}
