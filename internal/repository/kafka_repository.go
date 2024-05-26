package repository

import (
	"github.com/segmentio/kafka-go"
)

type KafkaRepository struct {
	producer *kafka.Writer
}

func NewKafkaRepository(producer *kafka.Writer) *KafkaRepository {
	return &KafkaRepository{
		producer: producer,
	}
}
