include .env
GOOSE = goose
MIGRATIONS_DIR = ./cmd/migrate/migrations

.PHONY: goose-up
goose-up:
	$(GOOSE) postgres $(DB_URL) up -dir $(MIGRATIONS_DIR)

.PHONY: goose-down
goose-down:
	$(GOOSE) postgres $(DB_URL) down -dir $(MIGRATIONS_DIR)

.PHONY: goose-create
goose-create:
	@read -p "Enter migration name: " name; \
	$(GOOSE) -s create $name sql -dir $(MIGRATIONS_DIR)

