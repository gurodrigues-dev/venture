package service

import (
	"context"
	"gin/internal/repository"
	"gin/types"
	"gin/utils"
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

func (s *Service) CreateSchool(ctx context.Context, school *types.School) error {
	school.Password = utils.HashPassword(school.Password)
	return s.repository.CreateSchool(ctx, school)
}

func (s *Service) ReadSchool(ctx context.Context, id *int) (*types.School, error) {
	return s.repository.ReadSchool(ctx, id)
}

func (s *Service) UpdateCreateSchool(ctx context.Context) error {
	return s.repository.UpdateSchool(ctx)
}

func (s *Service) DeleteCreateSchool(ctx context.Context, id *int) error {
	return s.repository.DeleteSchool(ctx, id)
}

func (s *Service) AuthSchool(ctx context.Context, school *types.School) error {
	return s.repository.AuthSchool(ctx, school)
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
