package redis

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/redis/go-redis/v9"

	"github.com/vladimirdotk/news-bot/internal/domain"
)

type QueueService struct {
	redisClient *redis.Client
	log         *slog.Logger
}

func NewQueueService(redisClient *redis.Client, log *slog.Logger) *QueueService {
	return &QueueService{
		redisClient: redisClient,
		log: slog.With(
			slog.Group("queue service"),
		),
	}
}

func (q *QueueService) Publish(ctx context.Context, topic domain.QueueTopic, data interface{}) error {
	if err := q.redisClient.RPush(ctx, topic.String(), data).Err(); err != nil {
		return fmt.Errorf("rpush to redis: %v", err)
	}

	return nil
}
