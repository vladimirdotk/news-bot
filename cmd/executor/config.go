package main

import "github.com/vladimirdotk/news-bot/config"

// Config describes configuration for (command) executor.
type Config struct {
	Telegram config.Telegram
	Redis    config.Redis
}
