postgres:
	docker exec -it komek-db psql -U app

migrateup:
	migrate -source file://db/migrations -database "postgres://app:secret@localhost:5433/app?sslmode=disable" -verbose up

migratedown:
	migrate -source file://db/migrations -database "postgres://app:secret@localhost:5433/app?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

PHONY: postgres