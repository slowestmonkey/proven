run:
	go run .

migrate-up:
	migrate -source file://db/migrations -database "postgres://localhost:5432/proven?sslmode=disable" up

migrate-down:
	migrate -source file://db/migrations -database "postgres://localhost:5432/proven?sslmode=disable" down