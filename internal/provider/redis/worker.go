package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v7"

	"github.com/vladimirdotk/news-bot/internal/domain"
)

type Worker struct {
	redisClient     redis.Cmdable
	commandExecutor CommandExecutor
}

func NewWorker(redisClient redis.Cmdable, commandExecutor CommandExecutor) *Worker {
	return &Worker{
		redisClient:     redisClient,
		commandExecutor: commandExecutor,
	}
}

// Run executes worker.
func (w *Worker) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("Starting worker...")

	ticker := time.NewTicker(time.Second)

loop:
	for {
		select {
		case <-ctx.Done():
			log.Printf("Worker context is done")
			break loop
		case <-ticker.C:
			if err := w.execute(); err != nil {
				log.Printf("Worker execute: %v", err)
			}
		}
	}
}

func (w *Worker) execute() error {
	response, err := w.redisClient.BLPop(time.Second*2, domain.QueueTopicIncomingCommand.String()).Result()
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

	return w.commandExecutor.Exec(*incomingMessage)
}

func toIncomingMessage(src string) (*domain.IncomingMessage, error) {
	var incomingMessage domain.IncomingMessage
	if err := json.Unmarshal([]byte(src), &incomingMessage); err != nil {
		return nil, fmt.Errorf("convert from json to incmming message: %v", err)
	}

	return &incomingMessage, nil
}

func getMessage(response []string) (string, error) {
	if len(response) != 2 {
		return "", fmt.Errorf("get message from response: %v", response)
	}

	return response[1], nil
}
