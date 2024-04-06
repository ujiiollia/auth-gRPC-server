package main

import (
	"app/internal/app"
	"app/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
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
	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	go func() {
		application.GRPCServer.MustRun()
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	endSignal := <-stop
	log.Info("stopping application", slog.String("signal", endSignal.String()))

	application.GRPCServer.Stop()
	log.Info("Gracefully stopped")
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
