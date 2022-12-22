-include .env

GO ?= go
MIGRATE	?= migrate

POSTGRESQL_URL = ${DB_DRIVER}://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}
MIGRATIONS_DIR = db/migrations

start:
	@$(GO) run .

migrations-up:
	@$(MIGRATE) -database $(POSTGRESQL_URL) -path $(MIGRATIONS_DIR) up

migrations-down:
	@$(MIGRATE) -database $(POSTGRESQL_URL) -path $(MIGRATIONS_DIR) down
