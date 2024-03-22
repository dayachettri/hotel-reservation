include .env

build:
	@docker compose build

run: build
	@docker compose up

test:
	@go test -v ./...

migrate-up:
	@goose -dir migrations postgres postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable up

migrate-down:
	@goose -dir migrations postgres postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable down

migrate-redo:
	@goose -dir migrations postgres postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable redo

migrate-status:
	@goose -dir migrations postgres postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable status



