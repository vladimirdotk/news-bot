package main

import (
	"log"
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/vladimirdotk/news-bot/internal/logger"
	"github.com/vladimirdotk/news-bot/internal/news"

	"github.com/redis/go-redis/v9"
)

func main() {
	config := Config{}

	if err := cleanenv.ReadEnv(&config); err != nil {
		log.Fatalf("parsing config: %v", err)
	}

	logLevel := slog.Level(config.Log.Level)
	log := logger.AssembleLogger(logLevel, "collector")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.Db,
	})

	_ = news.NewCollector(redisClient, log)
}
