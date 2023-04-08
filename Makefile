run:

	docker compose --profile tools run --rm migrate up &
	docker compose up

run_local:
	go run cmd/main.go

local_migration_up:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up

local_migration_down:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down

.PHONY: run