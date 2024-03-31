package main

import "github.com/vladimirdotk/news-bot/config"

// Config describes configuration for news collector.
type Config struct {
	Redis config.Redis
	Log   config.Log
}
