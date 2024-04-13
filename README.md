# News bot

![build](https://github.com/vladimirdotk/news-bot/actions/workflows/golang-checks.yml/badge.svg)

> Under heavy development...

## Install dependencies

- `make deps`

## Build

- `make build`

## Dev Run

- build binaries (see above)
- run redis `docker-compose up -d`
- run bot `TELEGRAM_BOT_TOKEN="your-token" REDIS_ADDR="localhost:6379 ./bin/bot`
- run command executor `TELEGRAM_BOT_TOKEN="your-token" REDIS_ADDR="localhost:6379 ./bin/executor`
- run collector `REDIS_ADDR="localhost:6379 ./bin/collector`

## Tests

- run `make generate` to generate mocks (optional)
- run `make test`
