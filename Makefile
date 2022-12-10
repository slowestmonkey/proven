run:
	go run ./cmd/app

test:
	go test -v -cover ./...

watch:
	nodemon --exec go run ./app --signal SIGTERM

migrate-up:
	migrate -source file://migrations -database "postgres://localhost:5432/proven?sslmode=disable" up

migrate-down:
	migrate -source file://migrations -database "postgres://localhost:5432/proven?sslmode=disable" down

compose-up:
	docker-compose up --build -d postgres app && docker-compose logs -f

compose-down:
	docker-compose down --remove-orphans