build:
	@docker compose build

run: build
	@docker compose up -d

test:
	@go test -v ./...


