package main

import (
	"log"

	"github.com/vladimirdotk/news-bot/internal/news"

	"github.com/caarlos0/env/v6"
	"github.com/go-redis/redis/v7"
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
