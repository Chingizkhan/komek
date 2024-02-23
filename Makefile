postgres:
	docker exec -it komek-db psql -U app

migrate_up:
	migrate -source file://db/migrations -database "postgres://app:secret@localhost:5433/app?sslmode=disable" -verbose up

migrate_down:
	migrate -source file://db/migrations -database "postgres://app:secret@localhost:5433/app?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

app:
	go run cmd/app/app.go

mock_banking:
	mockgen -package mock_banking -destination internal/service/banking/mock/banking.go komek/internal/controller/http/v1 Banking

PHONY: postgres migrate_up migrate_down test app mock_banking