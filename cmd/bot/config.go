package main

import (
	"github.com/vladimirdotk/news-bot/config"
)

// Config describes configuration for bot.
type Config struct {
	Telegram config.Telegram
	Redis    config.Redis
	Log      config.Log
}
