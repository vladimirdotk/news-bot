package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/vladimirdotk/news-bot/internal/domain"
)

type Worker struct {
	redisClient     redis.Cmdable
	commandExecutor CommandExecutor
	log             *slog.Logger
}

func NewWorker(redisClient redis.Cmdable, commandExecutor CommandExecutor, log *slog.Logger) *Worker {
	return &Worker{
		redisClient:     redisClient,
		commandExecutor: commandExecutor,
		log: slog.With(
			slog.Group("redis worker"),
		),
	}
}

// Run executes worker.
func (w *Worker) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	w.log.Info("Starting worker...")

	ticker := time.NewTicker(time.Second)

loop:
	for {
		select {
		case <-ctx.Done():
			w.log.Info("Worker context is done")
			break loop
		case <-ticker.C:
			if err := w.execute(ctx); err != nil {
				w.log.Info("Worker execute: %v", err)
			}
		}
	}
}

func (w *Worker) execute(ctx context.Context) error {
	response, err := w.redisClient.
		BLPop(
			ctx, time.Second*2,
			domain.QueueTopicIncomingCommand.String(),
		).
		Result()
	if err == redis.Nil {
		return nil
	}
	if err != nil {
		return fmt.Errorf("blpop incoming message: %v", err)
	}

	message, err := getMessage(response)
	if err != nil {
		return fmt.Errorf("get message from redis response: %v", err)
	}

	incomingMessage, err := toIncomingMessage(message)
	if err != nil {
		return fmt.Errorf("to incoming message: %v", err)
	}

	return w.commandExecutor.Exec(ctx, incomingMessage)
}

func toIncomingMessage(src string) (domain.IncomingMessage, error) {
	var incomingMessage domain.IncomingMessage
	if err := json.Unmarshal([]byte(src), &incomingMessage); err != nil {
		return domain.IncomingMessage{}, fmt.Errorf("convert from json to incoming message: %v", err)
	}

	return incomingMessage, nil
}

func getMessage(response []string) (string, error) {
	if len(response) != 2 {
		return "", fmt.Errorf("get message from response: %v", response)
	}

	return response[1], nil
}
