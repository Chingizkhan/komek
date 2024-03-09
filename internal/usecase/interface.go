package usecase

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type (
	Transactional interface {
		Exec(ctx context.Context, fn func(tx pgx.Tx) error) error
	}
)
