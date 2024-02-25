package banking_uc

import (
	"context"
	"komek/db/sqlc"
)

type (
	Service struct {
		tx Store
	}

	Store interface {
		Exec(ctx context.Context, fn func(*sqlc.Queries) error) error
	}
)

var txKey = struct{}{}

func New(store Store) *Service {
	return &Service{
		store,
	}
}
