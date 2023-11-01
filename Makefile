ifeq ($(POSTGRES_SETUP_TEST),)
	POSTGRES_SETUP_TEST := user=test password=test dbname=test host=localhost port=5432 sslmode=disable
endif

INTERNAL_PKG_PATH=$(CURDIR)
MIGRATION_FOLDER=$(INTERNAL_PKG_PATH)/migrations

build:
	docker-compose build

up:
	docker-compose up -d postgres zookeeper kafka1 kafka2 kafka3

down:
	docker-compose down

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql

.PHONY: migration-up
migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: migration-down
migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down

.PHONY: unit-test
unit-test:
	go test ./...

.PHONY: integration-test
integration-test:
	go test -tags=integration ./...
	go test -tags=integration_handler ./...

.PHONY: test
test: unit-test integration-test

.PHONY: clean-db
clean-db: migration-down migration-up

.PHONY: run
run:
	go run ./cmd/service

.PHONY: run-all
run-all: up migration-up test run

