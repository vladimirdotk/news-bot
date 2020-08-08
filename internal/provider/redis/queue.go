package redis

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

type QueueService struct {
	redisClient *redis.Client
}

func NewQueueService(redisClient *redis.Client) *QueueService {
	return &QueueService{redisClient: redisClient}
}

func (q *QueueService) Publish(topic string, data interface{}) error {
	if err := q.redisClient.RPush(topic, data).Err(); err != nil {
		return fmt.Errorf("rpush to redis: %v", err)
	}

	return nil
}
