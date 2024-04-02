package repository

import (
	"context"
	"gin/config"

	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	producer *kafka.Writer
	consumer *kafka.Reader
}

func NewKafkaClient() (*Kafka, error) {
	conf := config.Get()

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  conf.Messaging.Brokers,
		Topic:    conf.Messaging.Topic,
		Balancer: &kafka.LeastBytes{},
	})

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   conf.Messaging.Brokers,
		Topic:     conf.Messaging.Topic,
		Partition: conf.Messaging.Partition,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})

	kafkaClient := &Kafka{
		producer: writer,
		consumer: reader,
	}

	return kafkaClient, nil
}

func (k *Kafka) Producer(ctx context.Context) {

}

func (k *Kafka) Consumer(ctx context.Context) {

}

func (k *Kafka) Publish(ctx context.Context) {

}

func (k *Kafka) Subscribe(ctx context.Context) {

}
