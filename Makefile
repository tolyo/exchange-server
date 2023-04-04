
.DEFAULT_GOAL := help
.PHONY: help

setup:
	npm install

build: ## Installs and compiles dependencies
	go build -v ./...
	go install github.com/pressly/goose/v3/cmd/goose@latest

run: ## Start dev mode
	go run main.go

test:
	go test ./... -v -cover -p 1

lint:
	@go fmt ./...
	@go vet ./...
	@staticcheck ./...


DB_DSN:=$$(yq e '.DB_DSN' ./pkg/conf/dev.yaml)
MIGRATE_OPTIONS=-allow-missing -dir="./pkg/db/migrations"

db-update: ## Migrate down on database
	goose -v $(MIGRATE_OPTIONS) postgres "$(DB_DSN)" up

db-downgrade: ## Migrate up on database
	echo "$(MIGRATE_OPTIONS)"
	goose -v $(MIGRATE_OPTIONS) postgres "$(DB_DSN)" reset

db-rebuild: ## Reset the database
	$(MAKE) db-downgrade 
	$(MAKE) db-update

build-api: ## Build OpenAPI
	node node_modules/swagger-cli/swagger-cli.js bundle --dereference \
 		-o pkg/static/docs/api/openapi.json \
 		-t json \
 		-r openapi/openapi.yaml

help:
	grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| sed -n 's/^\(.*\): \(.*\)##\(.*\)/\1\3/p' 