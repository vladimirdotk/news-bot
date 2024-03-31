package redis

import (
	"context"

	"github.com/vladimirdotk/news-bot/internal/domain"
)

//go:generate minimock -g -i github.com/vladimirdotk/news-bot/internal/provider/redis.CommandExecutor -o ./mocks -s "_mock.go"

// CommandExecutor describes a service that executes user's commands
// received from different message systems.
type CommandExecutor interface {
	// Exec executes user's command and returns an error if any.
	Exec(ctx context.Context, message domain.IncomingMessage) error
}
