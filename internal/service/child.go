package service

import (
	"context"
	"gin/internal/repository"
)

type ChildService struct {
	childrepository repository.ChildRepository
}

func NewChildService(repo repository.ChildRepository) *ChildService {
	return &ChildService{
		childrepository: repo,
	}
}

func (s *Service) CreateChild(ctx context.Context) {

}

func (s *Service) ReadChild(ctx context.Context) {

}

func (s *Service) UpdateChild(ctx context.Context) {

}

func (s *Service) DeleteChild(ctx context.Context) {

}
