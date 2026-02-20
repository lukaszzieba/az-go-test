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
	go build -o ./bin/go-fun ./cmd/api/main.go

build-az:
	GOOS=linux GOARCH=amd64 go build -o ./bin/go-fun ./cmd/api/main.go

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

# Minimal Docker commands
docker-build:
	docker build -t go-fun .

docker-run:
	docker run -d --name go-fun -p 8080:8080 --env-file .env go-fun

docker-stop:
	docker stop go-fun && docker rm go-fun

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

.PHONY: docker-build docker-run docker-stop docker-up docker-down docker-logs
