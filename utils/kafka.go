package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"gin/models"
	"net/http"
	"time"

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

func StarterKafkaConsumer() {
	url := "http://localhost:8080/api/v1/users/consumer"

	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-ticker.C:
			err := callURL(url)
			if err != nil {
				fmt.Printf("Erro ao chamar a URL: %v\n", err)
			}
		}
	}
}

func callURL(url string) error {
	fmt.Println("enviando req")
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("Status da resposta: %s\n", resp.Status)
	return nil
}
