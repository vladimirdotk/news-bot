package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/vladimirdotk/news-bot/internal/domain"
)

type QueueService struct {
	redisClient *redis.Client
}

func NewQueueService(redisClient *redis.Client) *QueueService {
	return &QueueService{redisClient: redisClient}
}

func (q *QueueService) Publish(topic domain.QueueTopic, data interface{}) error {
	ctx := context.TODO()
	if err := q.redisClient.RPush(ctx, topic.String(), data).Err(); err != nil {
		return fmt.Errorf("rpush to redis: %v", err)
	}

	return nil
}
