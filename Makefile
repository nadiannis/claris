help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

## run: run the cmd/api server
.PHONY: run
run:
	go run ./cmd/api

## build: build the cmd/api server into a single binary
.PHONY: build
build:
	go build -o ./bin/api ./cmd/api

## start: build the cmd/api server into a single binary & run it
.PHONY: start
start: build
	./bin/api
