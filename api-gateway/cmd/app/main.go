package main

import (
	"apigateway/internal/app"
	"apigateway/internal/config"
	"apigateway/logger"
)

func main() {
	cfg := config.New()
	log := logger.SetupLogger(cfg.LogLevel)
	application := app.NewApp(log, cfg)
	forever := make(chan bool)
	go application.HTTPApp.Start()
	<-forever
}
