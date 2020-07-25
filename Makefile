.PHONY: deps
deps:
	@go mod download
	@go mod vendor
	@go mod tidy

.PHONY: build
build:
	@go build -o ./bin/bot ./cmd/bot
	@go build -o ./bin/executor ./cmd/executor

.PHONY: run
run:
	./bin/bot &
	./bin/executor

.PHONY: clean
clean:
	@rm -fv ./bin/*
