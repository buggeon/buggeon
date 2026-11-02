.PHONY: test run install-deps build

test:
	gotestsum --format testdox --format-icons=hivis ./internal/services/

install-deps:
	go mod download
	go mod tidy

run:
	go run cmd/main.go

build:
	go build