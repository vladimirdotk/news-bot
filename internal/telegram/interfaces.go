package telegram

import "github.com/vladimirdotk/news-bot/internal/domain"

// MessageHandler is a service for handling user messages.
type MessageHandler interface {
	// Handle processes user message.
	Handle(message *domain.IncomingMessage) error
}
