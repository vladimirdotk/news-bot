package main

import (
	"log"

	"github.com/caarlos0/env/v10"

	"github.com/vladimirdotk/news-bot/internal/news"

	"github.com/redis/go-redis/v9"
)

func main() {
	config := Config{}

	if err := env.Parse(&config); err != nil {
		log.Fatalf("parsing collector config: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.Db,
	})

	_ = news.NewCollector(redisClient)
}
