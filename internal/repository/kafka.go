package repository

import (
	"context"
	"gin/config"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	producer *kafka.Writer
}

func NewKafkaClient() (*Kafka, error) {
	conf := config.Get()

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{conf.Messaging.Brokers},
		Topic:    conf.Messaging.Topic,
		Balancer: &kafka.LeastBytes{},
	})

	kafkaClient := &Kafka{
		producer: writer,
	}

	return kafkaClient, nil
}

func (k *Kafka) Producer(ctx context.Context, msg string) error {

	message := []byte(msg)

	err := k.producer.WriteMessages(context.Background(), kafka.Message{
		Key:     nil,
		Value:   message,
		Time:    time.Now(),
		Headers: nil,
	})

	if err != nil {
		log.Fatalf("error to writing message on Kafka: %v", err)
	}

	return err
}
