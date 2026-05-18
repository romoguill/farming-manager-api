.DEFAULT_GOAL := help

DATABASE_URL ?= postgres://farming:farming@localhost:5432/farming_manager?sslmode=disable
HTTP_ADDR      ?= :8080

export DATABASE_URL
export HTTP_ADDR

BINARY := bin/server
CMD    := ./cmd

.PHONY: help db-up db-down db-logs build run dev tidy docker-build docker-up docker-down

help:
	@echo "Targets:"
	@echo "  db-up         Start Postgres container"
	@echo "  db-down       Stop all compose services"
	@echo "  db-logs       Follow Postgres logs"
	@echo "  build         Compile binary to $(BINARY)"
	@echo "  run           Build and run $(BINARY)"
	@echo "  dev           Run API with go run (local dev)"
	@echo "  tidy          Run go mod tidy"
	@echo "  docker-build  Build API Docker image"
	@echo "  docker-up     Start Postgres + API containers"
	@echo "  docker-down   Stop full stack containers"

db-up:
	docker compose up -d postgres

db-down:
	docker compose down

db-logs:
	docker compose logs -f postgres

build:
	go build -o $(BINARY) $(CMD)

run: build
	./$(BINARY)

dev:
	go run $(CMD)

tidy:
	go mod tidy

docker-build:
	docker compose build api

docker-up:
	docker compose --profile full up -d

docker-down:
	docker compose --profile full down
