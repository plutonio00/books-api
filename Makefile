.PHONY: build
build:
	go build -o ./.bin/app ./cmd/app/main.go

.PHONY: exec
exec:
	./.bin/app

.DEFAULT_GOAL := build