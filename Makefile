
install:
	go mod download

build:
	go build -o bin/app ./cmd/api

run-api:
	go run ./cmd/api

test:
	go test ./... -v

lint:
	golangci-lint run ./...