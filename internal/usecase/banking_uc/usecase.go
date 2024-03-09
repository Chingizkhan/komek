package banking_uc

import (
	"context"
	"komek/db/sqlc"
)

type (
	UseCase struct {
		tx Store
	}

	Store interface {
		Exec(ctx context.Context, fn func(*sqlc.Queries) error) error
	}
)

var txKey = struct{}{}

func New(store Store) *UseCase {
	return &UseCase{
		store,
	}
}
