package handlers

import (
	"encoding/json"
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
	marshaledMessage, err := messageToJSON(message)
	if err != nil {
		return fmt.Errorf("convert message to JSON: %v", err)
	}

	if err := m.queueService.Publish("incoming_message", marshaledMessage); err != nil {
		return fmt.Errorf("publish incoming message: %v", err)
	}

	return nil
}

func messageToJSON(src *domain.IncomingMessage) ([]byte, error) {
	b, err := json.Marshal(src)
	if err != nil {
		return nil, fmt.Errorf("marshal message: %+v", err)
	}

	return b, nil
}
