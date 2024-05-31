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
	CreateChild(ctx context.Context, child *types.Child) error
	ReadChildren(ctx context.Context, cpf *string) ([]types.Child, error)
	UpdateChild(ctx context.Context, child *types.Child) error
	DeleteChild(ctx context.Context, rg *string) error
	CreateStudentAndSponsor(ctx context.Context, sponsor *types.Sponsor) error
	IsSponsor(ctx context.Context, rg *string) bool
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
	sqlQuery := `SELECT id, name, cpf, email, street, number, zip, complement FROM responsibles WHERE cpf = $1 LIMIT 1`
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
	_, err = tx.Exec("DELETE FROM responsibles WHERE cpf = $1", *cpf)
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

func (r *ResponsibleRepository) CreateChild(ctx context.Context, child *types.Child) error {
	responsibleQuery := `SELECT id, name, cpf, email, street, number, zip, complement FROM responsibles WHERE cpf = $1 LIMIT 1`
	var responsible types.Child
	err := r.db.QueryRow(responsibleQuery, child.Responsible.CPF).Scan(
		&responsible.Responsible.ID,
		&responsible.Responsible.Name,
		&responsible.Responsible.CPF,
		&responsible.Responsible.Email,
		&responsible.Responsible.Street,
		&responsible.Responsible.Number,
		&responsible.Responsible.ZIP,
		&responsible.Responsible.Complement,
	)
	if err != nil || err == sql.ErrNoRows {
		return err
	}
	sqlQuery := `INSERT INTO children (name, rg, shift, responsibles, street, number, complement, zip) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = r.db.Exec(sqlQuery, child.Name, child.RG, child.Shift, responsible.Responsible.CPF, responsible.Responsible.Street, responsible.Responsible.Number, responsible.Responsible.Complement, responsible.Responsible.ZIP)
	return err
}

func (r *ResponsibleRepository) ReadChildren(ctx context.Context, cpf *string) ([]types.Child, error) {

	sqlQuery := `SELECT name, rg, shift, street, number, complement, zip FROM children WHERE responsibles = $1`

	rows, err := r.db.Query(sqlQuery, *cpf)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var children []types.Child

	for rows.Next() {
		var child types.Child
		err := rows.Scan(&child.Name, &child.RG, &child.Shift, &child.Responsible.Street, &child.Responsible.Number, &child.Responsible.Complement, &child.Responsible.ZIP)
		if err != nil {
			return nil, err
		}
		children = append(children, child)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return children, nil
}

func (r *ResponsibleRepository) UpdateChild(ctx context.Context, child *types.Child) error {
	sqlQuery := `SELECT name, shift, driver, school WHERE rg = $1 LIMIT 1`
	var childFound types.Child
	err := r.db.QueryRow(sqlQuery, child.RG).Scan(
		&childFound.Name,
		&childFound.Shift,
		&childFound.Driver,
		&childFound.School,
	)
	if err != nil || err == sql.ErrNoRows {
		return err
	}

	if childFound.Name != child.Name {
		_, err := r.db.Exec("UPDATE children SET name = $1 WHERE rg = $2", child.Name, child.RG)
		if err != nil {
			return err
		}
	}

	if childFound.Shift != child.Shift {
		_, err := r.db.Exec("UPDATE children SET shift = $1 WHERE rg = $2", child.Shift, child.RG)
		if err != nil {
			return err
		}
	}

	if childFound.Driver != child.Driver {
		_, err := r.db.Exec("UPDATE children SET driver = $1 WHERE rg = $2", child.Driver, child.RG)
		if err != nil {
			return err
		}
	}

	if childFound.School != child.School {
		_, err := r.db.Exec("UPDATE children SET school = $1 WHERE rg = $2", child.School, child.RG)
		if err != nil {
			return err
		}
	}

	return err
}

func (r *ResponsibleRepository) DeleteChild(ctx context.Context, rg *string) error {
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
	_, err = tx.Exec("DELETE FROM children WHERE rg = $1", *rg)
	return err
}

func (r *ResponsibleRepository) IsSponsor(ctx context.Context, rg *string) bool {
	return false
}
