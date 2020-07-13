package handlers

import (
	"fmt"

	"github.com/vladimirdotk/news-bot/internal/domain"
)

type MessageHandler struct {
	queueService QueueService
}

func NewMessageHandler(queueService QueueService) *MessageHandler {
	return &MessageHandler{
		queueService: queueService,
	}
}

func (m *MessageHandler) Handle(message *domain.IncomingMessage) error {
	if err := m.queueService.Publish("incoming_message", message); err != nil {
		return fmt.Errorf("publish incoming message: %v", err)
	}

	return nil
}
