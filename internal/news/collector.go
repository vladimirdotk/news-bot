package news

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/vladimirdotk/news-bot/internal/domain"
)

type Collector struct {
	redisClient redis.Cmdable
	log         *slog.Logger
}

func NewCollector(redisClient redis.Cmdable, log *slog.Logger) *Collector {
	return &Collector{
		redisClient: redisClient,
		log: slog.With(
			slog.Group("news collector"),
		),
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
			if err := c.run(ctx); err != nil {
				c.log.Error(
					"collector run",
					slog.String("error", err.Error()),
				)
			}
		}
	}
}

func (c *Collector) run(ctx context.Context) error {
	tasks, err := c.tasks(ctx)
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

func (c *Collector) tasks(ctx context.Context) ([]domain.Task, error) {
	keys, err := c.redisClient.Keys(ctx, "*").Result()
	if err != nil {
		return nil, fmt.Errorf("get users and sources: %v", err)
	}

	tasks := make([]domain.Task, 0, len(keys))

	for i, key := range keys {
		userID, err := domain.ExtractUserFromKey(key)
		if err != nil {
			return nil, fmt.Errorf("extract user from key=%s: %v", key, err)
		}

		rawSource, err := c.redisClient.Get(ctx, key).Result()
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
