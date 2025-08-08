package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger() {
	opts := &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	Logger = slog.New(handler)
	slog.SetDefault(Logger)
}

func GetLogger() *slog.Logger {
	if Logger == nil {
		InitLogger()
	}
	return Logger
}
