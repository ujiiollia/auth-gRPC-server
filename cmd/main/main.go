package main

import (
	"app/internal/config"
	"log/slog"
	"os"
)

func main() {
	//инициализирован конфиг
	cfg := config.MustLoad()
	//инициализировать логгер
	log := setupLogger(cfg.Env)
	log.Info("start application",
		slog.String("env", cfg.Env),
		slog.Any("cfg", cfg),
	)
	//todo: инициализировать логику
	//todo: запустить сервер
}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
