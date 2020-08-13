package redis

import "github.com/vladimirdotk/news-bot/internal/domain"

// CommandExecutor describes a service that executes user's commands
// received from different message systems.
type CommandExecutor interface {
	// Exec executes user's command and returns an error if any.
	Exec(message domain.IncomingMessage) error
}
