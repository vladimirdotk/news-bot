package news

import (
	"context"
	"fmt"
	"log"
	"sync"
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
	tasks, err := c.tasks()
	if err != nil {
		return fmt.Errorf("get tasks: %v", err)
	}

	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)
		go c.runTask(&wg, task)
	}

	wg.Wait()

	return nil
}

func (c *Collector) tasks() ([]domain.Task, error) {
	keys, err := c.redisClient.Keys("*").Result()
	if err != nil {
		return nil, fmt.Errorf("get users and sources: %v", err)
	}

	tasks := make([]domain.Task, 0, len(keys))

	for i, key := range keys {
		userID, err := domain.ExtractUserFromKey(key)
		if err != nil {
			return nil, fmt.Errorf("extract user from key=%s: %v", key, err)
		}

		rawSource, err := c.redisClient.Get(key).Result()
		if err != nil {
			return nil, fmt.Errorf("get key=%s raw source: %v", key, err)
		}

		source, err := domain.SourceFromJSON(rawSource)
		if err != nil {
			return nil, fmt.Errorf("source from JSON: %v", err)
		}

		tasks[i] = domain.Task{
			UserID: userID,
			Source: *source,
		}
	}

	return tasks, nil
}

func (c *Collector) runTask(wg *sync.WaitGroup, task domain.Task) {
	defer wg.Done()

	// TODO: get news from source, set last seen field, set news to queue
}
