generate:
	protoc --proto_path=api/proto api/proto/*.proto --go_out=. --go-grpc_out=.

run:
	go run ./cmd/main/main.go

build:
	go build -o ./build/ ./cmd/main/...