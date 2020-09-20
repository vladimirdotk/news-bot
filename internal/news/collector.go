package news

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/vladimirdotk/news-bot/internal/domain"
)

type Collector struct {
	redisClient redis.Cmdable
}

func NewCollector(redisClient redis.Cmdable) *Collector {
	return &Collector{
		redisClient: redisClient,
	}
}

func (c *Collector) Run(ctx context.Context) {
	ticker := time.NewTicker(time.Minute * 10)

loop:
	for {
		select {
		case <-ctx.Done():
			break loop
		case <-ticker.C:
			if err := c.run(); err != nil {
				log.Printf("collector run: %v", err)
			}
		}
	}
}

func (c *Collector) run() error {
	tasks, err := c.getTasks()
	if err != nil {
		return fmt.Errorf("get tasks: %v", tasks)
	}
	// TODO: collect news from tasks in goroutines, set news to queue
	return nil
}

func (c *Collector) getTasks() ([]domain.Task, error) {
	userKeys := domain.UserSourcesKey("*")

	keys, err := c.redisClient.Keys(userKeys).Result()
	if err != nil {
		return nil, fmt.Errorf("get users and sources: %v", err)
	}

	tasks := make([]domain.Task, 0, len(keys))

	for i, key := range keys {
		userID, err := domain.ExtractUserFromKey(key)
		if err != nil {
			return nil, fmt.Errorf("extract user from key=%s: %v", key, err)
		}

		rawSources, err := c.redisClient.SMembers(key).Result()
		if err != nil {
			return nil, fmt.Errorf("get sources: %v", err)
		}

		for _, rawSource := range rawSources {
			source, err := domain.SourceFromJSON(rawSource)
			if err != nil {
				return nil, fmt.Errorf("source from JSON: %v", err)
			}
			tasks[i] = domain.Task{
				UserID: userID,
				Source: source,
			}
		}
	}

	return tasks, nil
}
