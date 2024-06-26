package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/caarlos0/env/v10"

	"github.com/redis/go-redis/v9"

	"github.com/vladimirdotk/news-bot/internal/command"
	"github.com/vladimirdotk/news-bot/internal/logger"
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

	logLevel := slog.Level(config.Log.Level)
	log := logger.AssembleLogger(logLevel, "executor")

	messageSender := telegram.NewSender(&http.Client{}, config.Telegram.BotToken, log)
	sourceDetector := source.NewDetector(log)
	commandExecutor := command.NewExecutor(redisClient, messageSender, sourceDetector, log)
	worker := redisprovider.NewWorker(redisClient, commandExecutor, log)

	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}

	wg.Add(1)
	go worker.Run(ctx, &wg)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	<-sigint

	cancel()

	wg.Wait()
}
