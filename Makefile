.PHONY: deps
deps:
	@go mod download
	@go mod vendor
	@go mod tidy

.PHONY: build
build:
	@go build -o ./bin/bot ./cmd/bot
	@go build -o ./bin/executor ./cmd/executor
	@go build -o ./bin/collector ./cmd/collector

.PHONY: clean
clean:
	@rm -fv ./bin/*

.PHONY: generate
generate: tools
	@export PATH=$(shell pwd)/bin:$(PATH); go generate ./...

.PHONY: tools
tools: deps
	@go install github.com/gojuno/minimock/v3/cmd/minimock

.PHONY: test
test:
	@go test ./...
