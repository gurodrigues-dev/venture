package service

import (
	"context"
	"gin/internal/repository"
)

type KafkaService struct {
	kafkarepository repository.KafkaRepository
}

func NewKafkaService(kafka repository.KafkaRepository) *KafkaService {
	return &KafkaService{
		kafkarepository: kafka,
	}
}

func (ks *KafkaService) AddMessageInQueue(ctx context.Context, msg string) error {
	return ks.kafkarepository.Producer(ctx, msg)
}
