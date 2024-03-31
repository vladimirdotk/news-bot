package config

import (
	"errors"
	"log/slog"
	"strings"
)

type Log struct {
	Level LogLevel `env:"LOG_LEVEL" env-default:"INFO"`
}

type LogLevel slog.Level

func (ll *LogLevel) SetValue(level string) error {
	switch strings.ToUpper(level) {
	case "DEBUG":
		*ll = LogLevel(slog.LevelDebug)
	case "INFO":
		*ll = LogLevel(slog.LevelInfo)
	case "WARN":
		*ll = LogLevel(slog.LevelWarn)
	case "ERROR":
		*ll = LogLevel(slog.LevelError)
	default:
		return errors.New("unknown level type")
	}

	return nil
}
