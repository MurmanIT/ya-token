package main

import (
	"log/slog"
	"token/internal/config"
	"token/internal/transport/rest"
	"token/pkg/logger"
)

func main() {
	cfg := config.LoadConfig()
	logger := logger.Init(cfg)
	logger = logger.With(
		slog.String("env", cfg.Env),
		slog.Int("port", cfg.HttpServer.Port),
	)
	server := rest.Init(cfg, logger)
	logger.Info("Server is running..")
	server.Run()
}
