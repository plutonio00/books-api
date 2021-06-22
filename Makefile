.PHONY: build
build:
	go build -o ./.bin/app ./cmd/app/main.go

exec:
	./.bin/app

swag:
	swag init -g internal/app/app.go

.DEFAULT_GOAL := build

