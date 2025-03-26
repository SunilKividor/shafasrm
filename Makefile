MIGRATIONS_DIR=internal/database/migrations

DB_USER ?= shafasrmadmin
DB_PASSWORD ?= shafasrm2723
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_NAME ?= shafasrm
DB_SSLMODE ?= disable

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

.PHONY: migrate create up down redo status run

create:
    @read -p "Enter migration name: " name;\
    goose -dir ${MIGRATIONS_DIR} create $$name sql

up:
    goose -dir ${MIGRATIONS_DIR} postgres "${DB_URL}" up

down:
    goose -dir ${MIGRATIONS_DIR} postgres "${DB_URL}" down

redo:
    goose -dir ${MIGRATIONS_DIR} postgres "${DB_URL}" redo

status:
    goose -dir ${MIGRATIONS_DIR} postgres "${DB_URL}" status

run: up
    go run cmd/app/main.go