.PHONY: build
build:
	go build -o ./.bin/app ./cmd/app/main.go

exec:
	./.bin/app

swag:
	swag init -g cmd/app/main.go

.DEFAULT_GOAL := build

