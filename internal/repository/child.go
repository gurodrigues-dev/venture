package repository

import (
	"context"
	"gin/types"

	_ "github.com/lib/pq"
)

type ChildRepository interface {
	CreateChild(ctx context.Context, child *types.Child, responsible *types.Responsible) error
	ReadChild(ctx context.Context, idResponsible *int) ([]types.Child, error)
	UpdateChild(ctx context.Context) error
	DeleteChild(ctx context.Context, idChild *int) error
}

func (p *Postgres) CreateChild(ctx context.Context, child *types.Child, responsible *types.Responsible) error {
	return nil
}

func (p *Postgres) ReadChild(ctx context.Context, id *int) ([]types.Child, error) {
	return nil, nil
}

func (p *Postgres) UpdateChild(ctx context.Context) error {
	return nil
}

func (p *Postgres) DeleteChild(ctx context.Context, idChild *int) error {
	return nil
}
