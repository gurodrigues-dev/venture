package utils

import (
	"context"
	"encoding/json"
	"gin/models"

	"github.com/segmentio/kafka-go"
)

func PostToKafkaToCreateUsers(userData *models.CreateUser) (bool, error) {

	kafkaConfig := models.KafkaConfig{
		BrokerAddress: "localhost:9092",
		Topic:         "kafka.create.users",
	}

	p := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaConfig.BrokerAddress},
		Topic:    kafkaConfig.Topic,
		Balancer: &kafka.Hash{},
	})

	defer p.Close()

	userBytes, err := json.Marshal(userData)

	if err != nil {
		return false, err
	}

	err = p.WriteMessages(context.Background(), kafka.Message{
		Value: userBytes,
	})

	return true, nil

}
