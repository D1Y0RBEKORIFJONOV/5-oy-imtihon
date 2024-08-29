package main

import (
	"ekzamen_5/booking-service/internal/config"
	"ekzamen_5/booking-service/internal/infastructure/kafka/consumer"
	"ekzamen_5/booking-service/logger"
	"log"
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
	log.Println("Stopping consumer")
}
