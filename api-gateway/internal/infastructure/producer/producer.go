package producer

import (
	"context"
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
)

func Producer(key []byte, req []byte, topicName string) error {
	client, err := kgo.NewClient(
		kgo.SeedBrokers("broker:29092"),
		kgo.AllowAutoTopicCreation(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx := context.Background()
	if err := client.Ping(ctx); err != nil {
		log.Println("client not connected to kafka", err)
	}

	record := kgo.Record{
		Key:   key,
		Topic: topicName,
		Value: req,
	}
	if err := client.ProduceSync(ctx, &record).FirstErr(); err != nil {
		log.Println(err)
	}
	return nil
}
