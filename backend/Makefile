migrate:
	atlas migrate apply --url "postgres://$(user):$(pwd)@localhost:5432/budget_tracker?sslmode=disable&search_path=public" --exec-order non-linear

server:
	go run cmd/server/main.go
