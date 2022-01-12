.PHONY: all
all: build
FORCE: ;

SHELL  := env LIBRARY_ENV=$(LIBRARY_ENV) $(SHELL)
LIBRARY_ENV ?= dev

BIN_DIR = $(PWD)/bin

.PHONY: build

build-api:
	go build -tags $(LIBRARY_ENV) -o ./bin/cmd cmd/main.go

dependencies:
	go mod download

build: dependencies build-api

ci: dependencies