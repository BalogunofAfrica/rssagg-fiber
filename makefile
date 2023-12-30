# Makefile

include .env
export

.PHONY: migrate-up
migrate-up:
	cd sql/schema && goose postgres "$(DB_URL)" up

.PHONY: migrate-down
migrate-down:
	cd sql/schema && goose postgres "$(DB_URL)" down

.PHONY: generate-client
generate-client:
	sqlc generate