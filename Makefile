# Makefile para app-challenge

.PHONY: build run test coverage clean

build:
	go build -o app-challenge ./cmd/main.go

run: build
	./app-challenge

test:
	go test -v ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
