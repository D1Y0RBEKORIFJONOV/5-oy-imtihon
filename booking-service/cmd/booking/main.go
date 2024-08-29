package main

import (
	"ekzamen_5/booking-service/internal/app"
	"ekzamen_5/booking-service/internal/config"
	"ekzamen_5/booking-service/logger"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.New()
	log := logger.SetupLogger(cfg.LogLevel)
	log.Info("Starting service1", slog.Any(
		"config", cfg.RPCPort))
	application := app.NewApp(log, cfg)

	go application.GrpcServer.Run()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sig := <-stop
	log.Info("received shutdown signal", slog.String("signal", sig.String()))
	application.GrpcServer.Stop()
	log.Info("shutting down server")
}
