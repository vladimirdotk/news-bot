package main

import (
	"context"
	"log"
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/go-redis/redis/v7"
	"github.com/vladimirdotk/news-bot/internal/command"
	"github.com/vladimirdotk/news-bot/internal/provider/redisserver"
	"github.com/vladimirdotk/news-bot/internal/provider/telegram"
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
	commandExecutor := command.NewExecutor(redisClient, messageSender)

	worker := redisserver.NewWorker(redisClient, commandExecutor)
	worker.Run(context.TODO())
}
