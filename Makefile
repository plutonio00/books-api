.PHONY: build
build:
	go build -o ./.bin/app ./cmd/app/main.go
	#go build -o ./.bin/migrations ./cmd/migrations/main.go

build_db_init_command:
	go build -o ./.bin/db_init ./cmd/db_init/main.go

exec:
	./.bin/app

swag:
	swag init -g cmd/app/main.go

db_init:
	./.bin/db_init

.DEFAULT_GOAL := build

