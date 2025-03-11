.PHONY: migrations-up migrations-drop

# Default values (can be overwritten by environment variables set in the script)
DB_USER ?= null
DB_PASSWORD ?= null
DB_HOST ?= null
DB_PORT ?= null
DB_NAME ?= null
SSL_MODE ?= null

MIGRATE_IMAGE = migrate/migrate:4
MIGRATION_PATH = $(PWD)/pkg/cipher/repository/postgres/migrations
NETWORK = --network=host
CONTAINER_NAME = go_db

DB_URL = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(SSL_MODE)


# Apply migrations
migrations-up:
	docker run --rm $(NETWORK) -v $(MIGRATION_PATH):/migrations \
		$(MIGRATE_IMAGE) -path=/migrations -database "$(DB_URL)" up || exit 1
	@echo "Migrations applied successfully"

# Drop and recreate the schema
migrations-drop:
	docker run --rm $(NETWORK) -v $(MIGRATION_PATH):/migrations \
		$(MIGRATE_IMAGE) -path=/migrations -database "$(DB_URL)" down || exit 1
	@echo "Schema dropped and recreated successfully"

print-vars:
	@echo "DB_USER = $(DB_USER)"
	@echo "DB_PASSWORD = $(DB_PASSWORD)"
	@echo "DB_HOST = $(DB_HOST)"
	@echo "DB_PORT = $(DB_PORT)"
	@echo "DB_NAME = $(DB_NAME)"
	@echo "SSL_MODE = $(SSL_MODE)"