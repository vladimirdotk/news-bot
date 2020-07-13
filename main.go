package main

import (
	"log"

	"github.com/vladimirdotk/news-bot/internal/provider/natsserver"

	"github.com/caarlos0/env/v6"
	nats "github.com/nats-io/nats.go"
	"github.com/vladimirdotk/news-bot/internal/handlers"
	"github.com/vladimirdotk/news-bot/internal/telegram"
)

type Config struct {
	TelegramBotToken string `env:"TELEGRAM_BOT_TOKEN,required"`
}

func main() {
	config := Config{}

	if err := env.Parse(&config); err != nil {
		log.Fatalf("parsing config: %v", err)
	}

	natsClient, err := nats.Connect(nats.DefaultURL, nats.NoReconnect())

	if err != nil {
		log.Fatalf("connecting to natsserver: %v", err)
	}

	natsEncodedConn, err := nats.NewEncodedConn(natsClient, nats.JSON_ENCODER)
	if err != nil {
		log.Fatalf("create encoded connection: %v", err)
	}
	defer natsEncodedConn.Close()

	//go:generate
	queueService := natsserver.NewQueueService(natsEncodedConn)
	messageHandler := handlers.NewMessageHandler(queueService)

	bot, err := telegram.NewBot(config.TelegramBotToken, messageHandler, true)
	if err != nil {
		log.Fatalf("create new bot: %v", err)
	}

	bot.Run()
}
