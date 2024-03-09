package banking_uc

import (
	"komek/db/sqlc"
	store2 "komek/internal/repos/tx"
	"komek/pkg/postgres"
	"log"
	"os"
	"testing"
)

const (
	dataSource = "postgresql://app:secret@localhost:5433/app?sslmode=disable"
)

var (
	testQueries *sqlc.Queries
	service     *UseCase
)

func TestMain(m *testing.M) {
	pg, err := postgres.New(
		dataSource,
		postgres.MaxPoolSize(20),
	)
	if err != nil {
		log.Fatal("can not connect postgres:", err)
	}
	testQueries = sqlc.New(pg.Pool)
	store := store2.NewTX(pg)
	service = New(store)

	os.Exit(m.Run())
}
