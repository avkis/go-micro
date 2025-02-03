.DEFAULT_GOAL := build

.PHONY:fmt vet build
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build ./cmd/web

## start: starts the application
start:
	@echo "Starting front-end"
	go run ./cmd/web
