package command

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/vladimirdotk/news-bot/internal/domain"
)

type Handler struct {
	queueService QueueService
	log          *slog.Logger
}

func NewHandler(queueService QueueService, log *slog.Logger) *Handler {
	return &Handler{
		queueService: queueService,
		log: slog.With(
			slog.Group("command handler"),
		),
	}
}

func (h *Handler) Handle(ctx context.Context, command *domain.IncomingMessage) error {
	marshaledCommand, err := commandToJSON(command)
	if err != nil {
		return fmt.Errorf("convert command to JSON: %v", err)
	}

	if err := h.queueService.Publish(ctx, domain.QueueTopicIncomingCommand, marshaledCommand); err != nil {
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
