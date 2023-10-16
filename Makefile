-include .env
.SILENT:
CURRENT_DIR=$(shell pwd)
DB_URL="postgres://otajonov:quvonchbek@localhost:5432/clinic?sslmode=disable"


run:
	go run cmd/main.go

migrate-up:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" -verbose down

migrate_file:
	migrate create -ext sql -dir migrations/ -seq alter_some_table

local-up:
	docker compose --env-file ./.env.docker up -d

pull-sub-module:
	git submodule update --init --recursive

update-sub-module:
	git submodule update --remote --merge

swag-gen:
	swag init -g api/router.go -o api/docs

.PHONY: run migrateup migratedown local-up proto-gen pull-sub-module update-sub-module