package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gin/types"

	_ "github.com/lib/pq"
)

type ResponsibleRepository interface {
	CreateResponsible(ctx context.Context, responsible *types.Responsible) error
	ReadResponsible(ctx context.Context, cpf *string) (*types.Responsible, error)
	UpdateResponsible(ctx context.Context) error
	DeleteResponsible(ctx context.Context, cpf *string) error
	AuthResponsible(ctx context.Context, responsible *types.Responsible) (*types.Responsible, error)
}

func (p *Postgres) CreateResponsible(ctx context.Context, responsible *types.Responsible) error {
	sqlQuery := `INSERT INTO responsibles (name, cpf, email, password, street, number, complement, zip) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := p.conn.Exec(sqlQuery, responsible.Name, responsible.CPF, responsible.Email, responsible.Password, responsible.Street, responsible.Number, responsible.Complement, responsible.ZIP)
	return err
}

func (p *Postgres) ReadResponsible(ctx context.Context, cpf *string) (*types.Responsible, error) {
	sqlQuery := `SELECT id, name, cpf, email, street, number, zip, complement FROM responsibles WHERE cnh = $1 LIMIT 1`
	var responsbile types.Responsible
	err := p.conn.QueryRow(sqlQuery, *cpf).Scan(
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

func (p *Postgres) UpdateResponsible(ctx context.Context) error {
	return nil
}

func (p *Postgres) DeleteResponsible(ctx context.Context, cpf *string) error {
	tx, err := p.conn.Begin()
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

func (p *Postgres) AuthResponsible(ctx context.Context, responsible *types.Responsible) (*types.Responsible, error) {
	sqlQuery := `SELECT id, name, cpf, email, password FROM responsibles WHERE email = $1 LIMIT 1`
	var responsibleData types.Responsible
	err := p.conn.QueryRow(sqlQuery, responsible.Email).Scan(
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
