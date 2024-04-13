package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/redis/go-redis/v9"
	"github.com/vladimirdotk/news-bot/internal/command"
	"github.com/vladimirdotk/news-bot/internal/logger"
	redisprovider "github.com/vladimirdotk/news-bot/internal/provider/redis"
	"github.com/vladimirdotk/news-bot/internal/telegram"
)

func main() {
	config := Config{}

	if err := cleanenv.ReadEnv(&config); err != nil {
		log.Fatalf("parsing config: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.Db,
	})

	logLevel := slog.Level(config.Log.Level)
	log := logger.AssembleLogger(logLevel, "telegram bot")

	queueService := redisprovider.NewQueueService(redisClient, log)
	commandHandler := command.NewHandler(queueService, log)

	bot, err := telegram.NewBot(config.Telegram.BotToken, commandHandler, true, log)
	if err != nil {
		log.Error(
			"create bot",
			slog.String("error", err.Error()),
		)
		os.Exit(1)
	}

	bot.Run()
}
