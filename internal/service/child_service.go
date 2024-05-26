package service

import "gin/internal/repository"

type ChildService struct {
	childrepository repository.ChildRepositoryInterface
}

func NewChildService(repo repository.ChildRepositoryInterface) *ChildService {
	return &ChildService{childrepository: repo}
}
