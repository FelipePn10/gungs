include .env

PWD := $(shell pwd)

MIGRATIONS_DIR := $(PWD)/internal/database/migrations

DATABASE_URL := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

create_migration:
	migrate create -ext=sql -dir=$(MIGRATIONS_DIR) -seq init

migrate_up:
	migrate -path=$(MIGRATIONS_DIR) \
		-database "$(DATABASE_URL)" \
		-verbose up

migrate_down:
	migrate -path=$(MIGRATIONS_DIR) \
		-database "$(DATABASE_URL)" \
		-verbose down

migrate_force:
	@if [ -z "$(v)" ]; then \
		echo "Uso: make migrate_force v=<version>"; \
		exit 1; \
	fi
	migrate -path=$(MIGRATIONS_DIR) \
		-database "$(DATABASE_URL)" \
		force $(v)

reset:
	migrate -path=$(MIGRATIONS_DIR) \
	    -database "$(DATABASE_URL)" \
	    drop -f

.PHONY: create_migration migrate_up migrate_down migrate_force reset