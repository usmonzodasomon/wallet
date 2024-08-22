run:
	go run cmd/wallet/main.go

migrate_up:
	goose -dir migrations postgres "postgresql://postgres:password@127.0.0.1:5432/wallet?sslmode=disable" up

migrate_down:
	goose -dir migrations postgres "postgresql://postgres:password@127.0.0.1:5432/wallet?sslmode=disable" down