package main

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/go-redis/redis/v7"
	"github.com/vladimirdotk/news-bot/internal/handlers"
	"github.com/vladimirdotk/news-bot/internal/provider/redisserver"
	"github.com/vladimirdotk/news-bot/internal/telegram"
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

	queueService := redisserver.NewQueueService(redisClient)
	messageHandler := handlers.NewMessageHandler(queueService)

	bot, err := telegram.NewBot(config.Telegram.BotToken, messageHandler, true)
	if err != nil {
		log.Fatalf("create new bot: %v", err)
	}

	bot.Run()
}
