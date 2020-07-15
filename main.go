package main

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/go-redis/redis/v7"
	"github.com/vladimirdotk/news-bot/internal/handlers"
	"github.com/vladimirdotk/news-bot/internal/provider/redisserver"
	"github.com/vladimirdotk/news-bot/internal/telegram"
)

type Config struct {
	TelegramBotToken string `env:"TELEGRAM_BOT_TOKEN,required"`
	Redis            RedisConfig
}

type RedisConfig struct {
	Addr     string `env:"REDIS_ADDR,required"`
	Password string `env:"REDIS_PASSWORD"`
	Db       int    `env:"REDIS_DB"`
}

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

	bot, err := telegram.NewBot(config.TelegramBotToken, messageHandler, true)
	if err != nil {
		log.Fatalf("create new bot: %v", err)
	}

	bot.Run()
}
