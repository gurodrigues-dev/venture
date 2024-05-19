package service

import (
	"context"
	"gin/internal/repository"
	"gin/types"
	"gin/utils"
)

type ResponsbileService struct {
	responsiblerepository repository.ResponsibleRepository
}

func NewResponsibleService(repo repository.ResponsibleRepository) *ResponsbileService {
	return &ResponsbileService{
		responsiblerepository: repo,
	}
}

func (rs *ResponsbileService) CreateResponsible(ctx context.Context, responsbile *types.Responsible) error {
	responsbile.Password = utils.HashPassword(responsbile.Password)
	return rs.responsiblerepository.CreateResponsible(ctx, responsbile)
}

func (rs *ResponsbileService) ReadResponsible(ctx context.Context, cpf *string) (*types.Responsible, error) {
	return rs.responsiblerepository.ReadResponsible(ctx, cpf)
}

func (rs *ResponsbileService) UpdateResponsible(ctx context.Context) error {
	return rs.responsiblerepository.UpdateResponsible(ctx)
}

func (rs *ResponsbileService) DeleteResponsible(ctx context.Context, cpf *string) error {
	return rs.responsiblerepository.DeleteResponsible(ctx, cpf)
}

func (rs *ResponsbileService) AuthResponsible(ctx context.Context, responsible *types.Responsible) (*types.Responsible, error) {
	return nil, nil
}
