package models

type KafkaConfig struct {
	Topic         string
	BrokerAddress string
	GroupID       string
}
