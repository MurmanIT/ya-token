package logger

import (
	"log/slog"
	"os"
	"token/internal/config"
)

const (
	local = "local"
	dev   = "dev"
	prod  = "prod"
)

func Init(cfg *config.Config) *slog.Logger {
	var log *slog.Logger

	switch cfg.Env {
	case local:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)
		log.Info("Configured logger for local...", slog.Any("config:", cfg))
	case dev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)
		log.Info("Configured logger for dev...", slog.Any("config:", cfg))
	case prod:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}),
		)
	}
	return log
}
