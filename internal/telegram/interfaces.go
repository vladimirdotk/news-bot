package telegram

import "github.com/vladimirdotk/news-bot/internal/domain"

//go:generate minimock -g -i github.com/vladimirdotk/news-bot/internal/telegram.MessageHandler -o ./mocks -s "_mock.go"

// MessageHandler is a service for handling user messages.
type MessageHandler interface {
	// Handle processes user message.
	Handle(message *domain.IncomingMessage) error
}
