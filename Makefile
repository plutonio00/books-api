.PHONY: build
build:
	go build -o ./.bin/app ./cmd/app/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: exec
exec:
	./.bin/app

.DEFAULT_GOAL := build