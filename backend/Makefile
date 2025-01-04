include .env
export

migrate-hash:
	atlas migrate hash

migrate-diff:
	atlas migrate diff --env gorm

migrate:
	atlas migrate apply --url "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/budget_tracker?sslmode=disable&search_path=public" --exec-order non-linear

server:
	go run commands/server/main.go
