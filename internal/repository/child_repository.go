package repository

import "database/sql"

type ChildRepositoryInterface interface {
}

type ChildRepository struct {
	db *sql.DB
}

func NewChildRepository(db *sql.DB) *ChildRepository {
	return &ChildRepository{
		db: db,
	}
}
