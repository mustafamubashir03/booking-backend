MIGRATIONS_FOLDER=db/migrations
DB_URL=root:root123@(localhost:3306)/auth_db_dev

migrate-create:
	goose -dir $(MIGRATIONS_FOLDER) create $(name) sql

migrate-up:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" up

migrate-down:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" down