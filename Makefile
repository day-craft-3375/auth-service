include .env

# миграции
migration_dir=migrations

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make migrate-create name=<migration_name>. Example: make migrate-create name=create_users_table"; \
		exit 1; \
	fi
	migrate create -ext sql -dir $(migration_dir) -seq $(name);

migrate-up:
	migrate -database $(POSTGRES_URL) -path $(migration_dir) up

migrate-down:
	migrate -database $(POSTGRES_URL) -path $(migration_dir) down