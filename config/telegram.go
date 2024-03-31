package config

// Telegram configures telegram options.
type Telegram struct {
	// BotToken sets token for bot given by BotFather
	BotToken string `env:"TELEGRAM_BOT_TOKEN" env-required:"true"`
}
