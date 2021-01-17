package redis

import (
	"fmt"

	"github.com/go-redis/redis/v7"

	"github.com/vladimirdotk/news-bot/internal/domain"
)

type QueueService struct {
	redisClient *redis.Client
}

func NewQueueService(redisClient *redis.Client) *QueueService {
	return &QueueService{redisClient: redisClient}
}

func (q *QueueService) Publish(topic domain.QueueTopic, data interface{}) error {
	if err := q.redisClient.RPush(topic.String(), data).Err(); err != nil {
		return fmt.Errorf("rpush to redis: %v", err)
	}

	return nil
}
