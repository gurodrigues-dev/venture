package kafka

import (
	"context"
	"encoding/json"
	"gin/models"

	"github.com/segmentio/kafka-go"
)

func PostToKafkaToCreateUsers(userData *models.CreateUser, kafkaConfig *models.KafkaConfig) (bool, error) {

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

func kafkaConsumerToCreateUsers(kafkaConfig *models.KafkaConfig) {

}
