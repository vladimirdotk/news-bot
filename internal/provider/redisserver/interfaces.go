package redisserver

import "github.com/vladimirdotk/news-bot/internal/domain"

type CommandExecutor interface {
	Exec(message domain.IncomingMessage) error
}
