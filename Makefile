DB_URL = "postgres://app:secret@localhost:5433/app?sslmode=disable"

postgres:
	docker exec -it komek-db psql -U app

migrate_up:
	migrate -source file://db/migrations -database "$(DB_URL)" -verbose up

migrate_down:
	migrate -source file://db/migrations -database "$(DB_URL)" -verbose down

test:
	go test -v -cover ./...

app:
	go run cmd/app/app.go

mock_banking:
	mockgen -package mock_banking -destination internal/service/banking/mock/banking.go komek/internal/controller/http/v1 Banking

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
        proto/*.proto

.PHONY: postgres migrate_up migrate_down test app mock_banking proto