package logger

import (
	"log/slog"
	"os"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
)

func AssembleLogger(logLevel slog.Level, serviceName string) *slog.Logger {
	writer := os.Stdout

	var loggerHandler slog.Handler = slog.NewJSONHandler(
		writer,
		&slog.HandlerOptions{
			AddSource: logLevel == slog.LevelDebug,
			Level:     logLevel,
		},
	)
	if isatty.IsTerminal(writer.Fd()) {
		loggerHandler = tint.NewHandler(
			writer,
			&tint.Options{
				AddSource: logLevel == slog.LevelDebug,
			},
		)
	}

	return slog.
		New(loggerHandler).
		With(
			slog.Group("service", slog.String("name", serviceName)),
		)
}
