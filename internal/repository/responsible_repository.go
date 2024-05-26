package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gin/types"
)

type ResponsibleRepositoryInterface interface {
	CreateResponsible(ctx context.Context, responsible *types.Responsible) error
	ReadResponsible(ctx context.Context, cpf *string) (*types.Responsible, error)
	UpdateResponsible(ctx context.Context) error
	DeleteResponsible(ctx context.Context, cpf *string) error
	AuthResponsible(ctx context.Context, responsible *types.Responsible) (*types.Responsible, error)
}

type ResponsibleRepository struct {
	db *sql.DB
}

func NewResponsibleRepository(db *sql.DB) *ResponsibleRepository {
	return &ResponsibleRepository{
		db: db,
	}
}

func (r *ResponsibleRepository) CreateResponsible(ctx context.Context, responsible *types.Responsible) error {
	sqlQuery := `INSERT INTO responsibles (name, cpf, email, password, street, number, complement, zip) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := r.db.Exec(sqlQuery, responsible.Name, responsible.CPF, responsible.Email, responsible.Password, responsible.Street, responsible.Number, responsible.Complement, responsible.ZIP)
	return err
}

func (r *ResponsibleRepository) ReadResponsible(ctx context.Context, cpf *string) (*types.Responsible, error) {
	sqlQuery := `SELECT id, name, cpf, email, street, number, zip, complement FROM responsibles WHERE cnh = $1 LIMIT 1`
	var responsbile types.Responsible
	err := r.db.QueryRow(sqlQuery, *cpf).Scan(
		&responsbile.ID,
		&responsbile.Name,
		&responsbile.CPF,
		&responsbile.Email,
		&responsbile.Street,
		&responsbile.Number,
		&responsbile.ZIP,
		&responsbile.Complement,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	return &responsbile, nil
}

func (r *ResponsibleRepository) UpdateResponsible(ctx context.Context) error {
	return nil
}

func (r *ResponsibleRepository) DeleteResponsible(ctx context.Context, cpf *string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	_, err = tx.Exec("DELETE FROM responsibles WHERE cnh = $1", *cpf)
	return err
}

func (r *ResponsibleRepository) AuthResponsible(ctx context.Context, responsible *types.Responsible) (*types.Responsible, error) {
	sqlQuery := `SELECT id, name, cpf, email, password FROM responsibles WHERE email = $1 LIMIT 1`
	var responsibleData types.Responsible
	err := r.db.QueryRow(sqlQuery, responsible.Email).Scan(
		&responsibleData.ID,
		&responsibleData.Name,
		&responsibleData.CPF,
		&responsibleData.Email,
		&responsibleData.Password,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	match := responsibleData.Password == responsible.Password
	if !match {
		return nil, fmt.Errorf("email or password wrong")
	}
	responsibleData.Password = ""
	return &responsibleData, nil
}
