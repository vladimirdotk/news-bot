package telegram

import "github.com/vladimirdotk/news-bot/internal/domain"

type MessageHandler interface {
	Handle(message *domain.IncomingMessage) error
}
