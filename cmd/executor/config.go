package main

type Config struct {
	Telegram Telegram
	Redis    Redis
}

type Telegram struct {
	BotToken string `env:"TELEGRAM_BOT_TOKEN,required"`
}

type Redis struct {
	Addr     string `env:"REDIS_ADDR,required"`
	Password string `env:"REDIS_PASSWORD"`
	Db       int    `env:"REDIS_DB"`
}
