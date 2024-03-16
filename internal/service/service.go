package service

import (
	"context"
	"gin/internal/repository"
)

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

func (s *Service) CreateUser(ctx context.Context) {

}

func (s *Service) ReadUser(ctx context.Context) {

}

func (s *Service) UpdateUser(ctx context.Context) {

}

func (s *Service) DeleteUser(ctx context.Context) {

}

func (s *Service) CreateChild(ctx context.Context) {

}

func (s *Service) ReadChild(ctx context.Context) {

}

func (s *Service) UpdateChild(ctx context.Context) {

}

func (s *Service) DeleteChild(ctx context.Context) {

}

func (s *Service) CreateDriver(ctx context.Context) {

}

func (s *Service) ReadDriver(ctx context.Context) {

}

func (s *Service) UpdateDriver(ctx context.Context) {

}
func (s *Service) DeleteDriver(ctx context.Context) {

}

func (s *Service) CreateSchool(ctx context.Context) {

}

func (s *Service) ReadScCreateSchool(ctx context.Context) {

}

func (s *Service) UpdateScCreateSchool(ctx context.Context) {

}

func (s *Service) DeleteScCreateSchool(ctx context.Context) {

}

func (s *Service) CheckEmail(ctx context.Context) {

}

func (s *Service) SendEmail(ctx context.Context) {

}

func (s *Service) SaveImageBucket(ctx context.Context) {

}

func (s *Service) SaveKeyAndValue(ctx context.Context) {

}

func (s *Service) VerifyToken(ctx context.Context) {

}
