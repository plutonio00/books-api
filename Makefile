.PHONY: build
build:
	go build -o ./.bin/app ./cmd/app/main.go

.PHONY: server_run
exec:
	./.bin/app

.DEFAULT_GOAL := build