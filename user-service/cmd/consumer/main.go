package main

import (
	"log"
	"user_service_smart_home/internal/config"
	"user_service_smart_home/internal/infastructure/kafka/consumer"
	"user_service_smart_home/logger"
)

func main() {
	log1 := logger.SetupLogger("local")
	cfg := config.New()
	c, err := consumer.NewConsumer(cfg, log1)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Connecting to RabbitMQ")
	log.Println("Starting consumer")
	c.Consume()
}
