package repository

import (
	"context"
	"gin/types"
)

type ResponsibleRepository interface {
	CreateResponsible(ctx context.Context, responsible *types.Responsible) error
	ReadResponsible(ctx context.Context, cpf *string) (*types.Responsible, error)
	UpdateResponsible(ctx context.Context) error
	DeleteResponsible(ctx context.Context, cpf *string) error
	AuthResponsible(ctx context.Context, responsible *types.Responsible) (*types.Responsible, error)
}
