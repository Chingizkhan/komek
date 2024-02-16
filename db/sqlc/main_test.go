package sqlc

import (
	"context"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	driverName = "postgres"
	dataSource = "postgresql://app:secret@localhost:5433/app?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	//conn, err := sql.Open(driverName, dataSource)
	//if err != nil {
	//	log.Fatal("can not connect to db:", err)
	//}

	poolConfig, err := pgxpool.ParseConfig(dataSource)
	if err != nil {
		log.Fatal("can not parse config:", err)
	}
	pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatal("can not connect config:", err)
	}

	testQueries = New(pool)

	os.Exit(m.Run())
}
