.PHONY: migrations-up migrations-drop print-vars

DB_USER ?= postgres
DB_PASSWORD ?= postgres
DB_HOST ?= localhost
DB_NAME ?= mydb
SSL_MODE ?= disable

MIGRATE_IMAGE = migrate/migrate:4
NETWORK = --network=host

SERVICES = cipher user-management

MIGRATION_PATH_cipher = $(PWD)/pkg/cipher/repository/postgres/migrations
DB_PORT_cipher ?= 5431
DB_URL_cipher = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT_cipher)/$(DB_NAME)?sslmode=$(SSL_MODE)

MIGRATION_PATH_user-management = $(PWD)/pkg/user-management/migrations
DB_PORT_user-management ?= 5432
DB_URL_user-management = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT_user-management)/$(DB_NAME)?sslmode=$(SSL_MODE)

migrations-up migrations-drop:
	@if [ -z "$(SERVICE)" ]; then \
		echo "Please specify a service, e.g., make $@ SERVICE=cipher"; \
		exit 1; \
	fi
	@$(MAKE) $@-$(SERVICE)

migrations-up-%:
	docker run --rm $(NETWORK) -v $(MIGRATION_PATH_$*):/migrations \
		$(MIGRATE_IMAGE) -path=/migrations -database "$(DB_URL_$*)" up || exit 1
	@echo "Migrations for $* applied successfully"

migrations-drop-%:
	docker run --rm $(NETWORK) -v $(MIGRATION_PATH_$*):/migrations \
		$(MIGRATE_IMAGE) -path=/migrations -database "$(DB_URL_$*)" down 1 || exit 1
	@echo "$* table dropped successfully"

print-vars:
	@for service in $(SERVICES); do \
		echo "DB_USER = $(DB_USER)"; \
		echo "DB_PASSWORD = $(DB_PASSWORD)"; \
		echo "DB_HOST = $(DB_HOST)"; \
		echo "DB_PORT_$$service = $(DB_PORT_$$service)"; \
		echo "DB_NAME = $(DB_NAME)"; \
		echo "SSL_MODE = $(SSL_MODE)"; \
	done