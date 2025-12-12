.PHONY: build run docker-up docker-down test lint

build:
	go build -o bin/weather ./cmd/server

run:
	go run ./cmd/server

test:
	go test ./...

docker-up:
	cd deploy && docker compose up --build

docker-down:
	cd deploy && docker compose down
