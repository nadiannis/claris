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

## audit: tidy dependencies, format code, vet code, & test code
.PHONY: audit
audit:
	@echo 'Tidying & verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...
