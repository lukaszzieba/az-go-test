# Load environment variables from .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

dev:
	air

test:
	go test -v ./...

run:
	go run ./cmd/api

build:
	go build -o ./bin/az-go-test ./cmd/api/main.go

# Goose migration targets
migration-create:
	@read -p "Enter migration name: " name; \
	cd ./sql/schema && goose postgres "$(DB_MIGRATION_URL)" create $$name sql

migration-up:
	cd ./sql/schema && goose postgres "$(DB_MIGRATION_URL)" up

migration-down:
	cd ./sql/schema && goose postgres "$(DB_MIGRATION_URL)" down

# SQLC code generation
generate:
	sqlc generate
