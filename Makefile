ifneq (,$(wildcard .env))
include .env
endif 
export

GOOSE=goose
MIGRATIONS=./migrations
DB_DRIVER=postgres
DB_URL?=$(DATABASE_URL)

.PHONY: apply-migrations remove-migrations make-migrations build start

apply-migrations:
	$(GOOSE) -dir $(MIGRATIONS) $(DB_DRIVER) "$(DB_URL)" up

remove-migrations:
	$(GOOSE) -dir $(MIGRATIONS) $(DB_DRIVER) "$(DB_URL)" down

make-migrations:
	$(GOOSE) -dir $(MIGRATIONS) create $(name) sql

build:
	go build -o bin/app cmd/server/main.go

start: build
	./bin/app