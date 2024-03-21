build:
	@docker compose build goapp
	
run: build
	@docker compose up -d

test:
	@go test -v ./...


