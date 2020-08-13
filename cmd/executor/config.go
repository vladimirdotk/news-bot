package main

// Config describes configuration for (command) executor.
type Config struct {
	Telegram Telegram
	Redis    Redis
}

// Telegram configures telegram options.
type Telegram struct {
	// BotToken sets token for bot given by BotFather
	BotToken string `env:"TELEGRAM_BOT_TOKEN,required"`
}

// Redis configures redis options.
type Redis struct {
	// Addr sets host(or ip) and optional port for redis connection.
	Addr string `env:"REDIS_ADDR,required"`
	// Password sets redis password, optional.
	Password string `env:"REDIS_PASSWORD"`
	// Db sets DB number, optional.
	Db int `env:"REDIS_DB"`
}
