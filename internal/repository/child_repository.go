package repository

import (
	"context"
	"database/sql"
	"gin/types"
)

type ChildRepositoryInterface interface {
	CreateChild(ctx context.Context, child *types.Child) error
	ReadChild(ctx context.Context, rg *string) (*types.Child, error)
	UpdateChild(ctx context.Context) error
	DeleteChild(ctx context.Context, rg *string) error
}

type ChildRepository struct {
	db *sql.DB
}

func NewChildRepository(db *sql.DB) *ChildRepository {
	return &ChildRepository{
		db: db,
	}
}
