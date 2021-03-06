.PHONY: cronparser
build:
	go build -o ./cronparser ./cmd/cronparser/main.go

.PHONY: test
test:
	go test -race -covermode=atomic ./...

.PHONY: lint
lint:
	golangci-lint run