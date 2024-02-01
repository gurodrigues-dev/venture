package kafka

import (
	"context"
	"encoding/json"
	"fmt"
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

func kafkaConsumerToCreateUsers(kafkaConfig *models.KafkaConfig) (bool, error) {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{kafkaConfig.BrokerAddress},
		Topic:     kafkaConfig.Topic,
		GroupID:   kafkaConfig.GroupID,
		MaxBytes:  10e6, //10 MB
		Partition: 0,
	})

	defer r.Close()

	for {

		m, err := r.ReadMessage(context.Background())

		if err != nil {
			return false, fmt.Errorf("Erro ao ler a mensagem: %v", err)
		}

		var userData models.CreateUser

		err = json.Unmarshal(m.Value, &userData)

		if err != nil {
			return false, fmt.Errorf("Erro ao fazer unmarshall da mensagem: %v", err)
		}

		fmt.Println(&userData)
		return true, nil

	}

}
