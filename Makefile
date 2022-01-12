build-api:
	go build -tags $(LIBRARY_ENV) -o ./bin/cmd cmd/main.go

dependencies:
	go mod download

build: dependencies build-api

ci: dependencies