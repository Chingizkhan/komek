package banking

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"komek/db/sqlc"
	store2 "komek/internal/repos/store"
	"log"
	"os"
	"testing"
)

const (
	dataSource = "postgresql://app:secret@localhost:5433/app?sslmode=disable"
)

var (
	testQueries *sqlc.Queries
	testPool    *pgxpool.Pool
	service     *Service
)

func TestMain(m *testing.M) {
	poolConfig, err := pgxpool.ParseConfig(dataSource)
	if err != nil {
		log.Fatal("can not parse config:", err)
	}
	testPool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatal("can not connect config:", err)
	}

	testQueries = sqlc.New(testPool)

	store := store2.NewStore(testPool)
	service = New(store)

	os.Exit(m.Run())
}
