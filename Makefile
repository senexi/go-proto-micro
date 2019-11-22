GO_CMD ?= go
VERSION ?= 0.0.1
APP_NAME ?= go-proto-micro

VERSION=$(shell git rev-parse --short HEAD)
BUILD=$(shell date +%FT%T%z)

LDFLAGS=-ldflags "-w -s -X main.VERSION="$(VERSION)" -X main.BUILD="$(BUILD)" -X main.NAME="$(APP_NAME)""

RED=\033[0;31m
GREEN=\033[0;32m

check-requirements:
	scripts/checkRequirements.sh

install-requirements:
	scripts/installRequirements.sh

generate:
	statik -f -src=./web -dest=./generated

format:
	@echo "\n$(GREEN)### formatting code\e[0m"
	$(GO_CMD) fmt ./...

build: generate
	@echo "\n$(GREEN)### compile $(APP_NAME)\e[0m"
	env GOOS=linux GOARCH=amd64 $(GO_CMD) build $(LDFLAGS) -o bin/$(APP_NAME)

compile:
	@echo "\n$(GREEN)### compile $(APP_NAME)\e[0m"
	env GOOS=linux GOARCH=amd64 $(GO_CMD) build $(LDFLAGS) -o bin/$(APP_NAME)

run: build
	@echo "\n$(GREEN)### run $(APP_NAME)\e[0m"
	bin/$(APP_NAME) serve

.PHONY: list
list:
	@$(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'
