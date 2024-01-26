package app

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func init() {
	m, err := migrate.New(
		"file://db/migrations",
		"postgres://app:secret@localhost:5460/app?sslmode=disable")
	if err != nil {
		log.Printf("Migrate: postgres is trying to connect: %s", err)
		return
	}

	err = m.Up()
	defer m.Close()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Printf("Migrate: no change")
			return
		}

		log.Fatalf("Migrate: up error: %s", err)
		return
	}

	log.Printf("Migrate: up success")
}
