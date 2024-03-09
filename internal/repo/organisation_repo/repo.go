package organisation_repo

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"komek/db/sqlc"
	"komek/pkg/postgres"
)

const (
	_defaultEntityCap = 50
)

type Repository struct {
	pool *pgxpool.Pool
	q    *sqlc.Queries
}

func New(pg *postgres.Postgres) *Repository {
	return &Repository{pg.Pool, sqlc.New(pg.Pool)}
}
