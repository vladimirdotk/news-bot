package main

import (
	"context"
	"log"
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/go-redis/redis/v7"
	"github.com/vladimirdotk/news-bot/internal/command"
	redisprovider "github.com/vladimirdotk/news-bot/internal/provider/redis"
	"github.com/vladimirdotk/news-bot/internal/provider/telegram"
	"github.com/vladimirdotk/news-bot/internal/source"
)

func main() {
	config := Config{}

	if err := env.Parse(&config); err != nil {
		log.Fatalf("parsing config: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.Db,
	})

	messageSender := telegram.NewSender(&http.Client{}, config.Telegram.BotToken)
	sourceDetector := source.NewDetector()
	commandExecutor := command.NewExecutor(redisClient, messageSender, sourceDetector)

	worker := redisprovider.NewWorker(redisClient, commandExecutor)
	worker.Run(context.TODO())
}
