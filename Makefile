.PHONY: build
build:
		go build -v ./cmd/apiserver

.PHONY: test
test:
		go test -v -race -timeout 30s ./...

.PHONY: migrate
migrate:
		psql -U sardor -h localhost -d gopher_school -f ./migrations/create_users_down.sql
		psql -U sardor -h localhost -d gopher_school -f ./migrations/create_users_up.sql

.DEFAULT_GOAL := build