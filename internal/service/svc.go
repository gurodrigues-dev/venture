package service

import (
	"encoding/json"
	"fmt"
	"gin/internal/repository"
	"gin/types"
)

type Service struct {
	repository repository.Repository
	cloud      repository.AWSRepository
	redis      repository.RedisRepository
	broker     repository.KafkaRepository
}

func NewService(repo repository.Repository, cloud repository.AWSRepository, redis repository.RedisRepository, broker repository.KafkaRepository) *Service {
	return &Service{
		repository: repo,
		cloud:      cloud,
		redis:      redis,
		broker:     broker,
	}
}

func (s *Service) InterfaceToString(value interface{}) (*string, error) {
	switch v := value.(type) {
	case string:
		return &v, nil
	default:
		return nil, fmt.Errorf("value isn't string")
	}
}

func (s *Service) EmailStructToJSON(email *types.Email) (string, error) {

	json, err := json.Marshal(email)

	if err != nil {
		return "", err
	}

	return string(json), nil

}
