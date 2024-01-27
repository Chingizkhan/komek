package app

import (
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//func init() {
//	m, err := migrate.New(
//		"file://db/migrations",
//		"postgres://app:secret@db:5432/app?sslmode=disable")
//	if err != nil {
//		log.Printf("Migrate: postgres is trying to connect: %s", err)
//		return
//	}
//
//	err = m.Up()
//	defer m.Close()
//	if err != nil {
//		if errors.Is(err, migrate.ErrNoChange) {
//			log.Printf("Migrate: no change")
//			return
//		}
//
//		log.Fatalf("Migrate: up error: %s", err)
//		return
//	}
//
//	log.Printf("Migrate: up success")
//}
