package command

import (
	"encoding/json"
	"fmt"

	"github.com/vladimirdotk/news-bot/internal/domain"
)

type Handler struct {
	queueService QueueService
}

func NewHandler(queueService QueueService) *Handler {
	return &Handler{
		queueService: queueService,
	}
}

func (h *Handler) Handle(command *domain.IncomingMessage) error {
	marshaledCommand, err := commandToJSON(command)
	if err != nil {
		return fmt.Errorf("convert command to JSON: %v", err)
	}

	if err := h.queueService.Publish(domain.QueueTopicIncomingCommand, marshaledCommand); err != nil {
		return fmt.Errorf("publish incoming command: %v", err)
	}

	return nil
}

func commandToJSON(src *domain.IncomingMessage) ([]byte, error) {
	b, err := json.Marshal(src)
	if err != nil {
		return nil, fmt.Errorf("marshal command: %v", err)
	}

	return b, nil
}
