package logger

import (
	"log/slog"
	"os"
)

// New: create new logger instance
func New(env string) Logger {
	logLevel := slog.LevelDebug
	if env != "dev" {
		logLevel = slog.LevelInfo
	}

	handler := slog.NewJSONHandler(
		os.Stdout, &slog.HandlerOptions{
			Level: logLevel,
		},
	)

	return NewAppLogger(slog.New(handler))
}
