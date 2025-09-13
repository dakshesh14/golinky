ifneq (,$(wildcard .env))
include .env
endif 
export

GOOSE=goose
MIGRATIONS=./migrations
DB_DRIVER=postgres
DB_URL?=$(DATABASE_URL)

.PHONY: apply-migrations remove-migrations make-migrations

apply-migrations:
	$(GOOSE) -dir $(MIGRATIONS) $(DB_DRIVER) "$(DB_URL)" up

remove-migrations:
	$(GOOSE) -dir $(MIGRATIONS) $(DB_DRIVER) "$(DB_URL)" down

make-migrations:
	$(GOOSE) -dir $(MIGRATIONS) create $(name) sql