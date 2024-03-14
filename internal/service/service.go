package service

import "gin/internal/repository"

type Service struct {
	repository repository.Repository
	cloud      repository.CloudRepository
	redis      repository.CacheRepository
}

func New(repo repository.Repository, cloud repository.CloudRepository, redis repository.CacheRepository) *Service {
	return &Service{
		repository: repo,
		cloud:      cloud,
		redis:      redis,
	}
}
