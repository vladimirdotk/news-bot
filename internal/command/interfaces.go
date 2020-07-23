package command

import "github.com/vladimirdotk/news-bot/internal/domain"

type ResponseSender interface {
	Send(message domain.OutgoingMessage) error
}