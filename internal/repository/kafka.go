package repository

import (
	"context"
	"fmt"
	"gin/config"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	producer *kafka.Writer
	consumer *kafka.Reader
}

func NewKafkaClient() (*Kafka, error) {
	conf := config.Get()

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{conf.Messaging.Brokers},
		Topic:    conf.Messaging.Topic,
		Balancer: &kafka.LeastBytes{},
	})

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{conf.Messaging.Brokers},
		Topic:     conf.Messaging.Topic,
		Partition: conf.Messaging.Partition,
		GroupID:   "reader.kafka.group",
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})

	kafkaClient := &Kafka{
		producer: writer,
		consumer: reader,
	}

	return kafkaClient, nil
}

func (k *Kafka) Producer(ctx context.Context, msg string) {

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

	log.Println("publish message on Kafka")

}

func (k *Kafka) Consumer(ctx context.Context) {

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			message, err := k.consumer.ReadMessage(context.Background())
			if err != nil {
				log.Fatalf("Erro ao ler mensagem do Kafka: %v", err)
			}
			fmt.Printf("Mensagem recebida do Kafka: %s\n", string(message.Value))
		}
	}()

	<-signals

}
