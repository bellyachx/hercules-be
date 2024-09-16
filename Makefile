include .env

generate:
	protoc --proto_path=api/proto api/proto/*.proto --go_out=. --go-grpc_out=.

run:
	go run ./cmd/main/main.go

build:
	go build -o ./build/ ./cmd/main/...

test:
	go test ./...

# Database migrations
.DB_URL=${DB_DRIVER}://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable
MIGRATIONS_DIR=migrations

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(.DB_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(.DB_URL)" down

migrate-create:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq -digits 5 $(name)