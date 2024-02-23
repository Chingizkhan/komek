package organisation_repo

import (
	"github.com/jackc/pgx/v4/pgxpool"
	org_db "komek/db/sqlc"
	"komek/pkg/postgres"
)

const (
	_defaultEntityCap = 50
)

type Repository struct {
	pool *pgxpool.Pool
	q    *org_db.Queries
}

func New(pg *postgres.Postgres) *Repository {
	return &Repository{pg.Pool, org_db.New(pg.Pool)}
}
