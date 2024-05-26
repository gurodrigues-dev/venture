package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gin/types"
	"log"
)

type SchoolRepositoryInterface interface {
	CreateSchool(ctx context.Context, school *types.School) error
	ReadSchool(ctx context.Context, cnpj *string) (*types.School, error)
	ReadAllSchools(ctx context.Context) ([]types.School, error)
	UpdateSchool(ctx context.Context) error
	DeleteSchool(ctx context.Context, cnpj *string) error
	AuthSchool(ctx context.Context, school *types.School) (*types.School, error)
	CreateInvite(ctx context.Context, invite *types.Invite) error
	IsEmployee(ctx context.Context, cnh *string) error
	GetEmployees(ctx context.Context, cnpj *string) ([]types.Driver, error)
	DeleteEmployee(ctx context.Context, record_id *int) error
}

type SchoolRepository struct {
	db *sql.DB
}

func NewSchoolRepository(db *sql.DB) *SchoolRepository {
	return &SchoolRepository{
		db: db,
	}
}

func (s *SchoolRepository) CreateSchool(ctx context.Context, school *types.School) error {
	sqlQuery := `INSERT INTO schools (name, cnpj, email, password, street, number, zip) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.Exec(sqlQuery, school.Name, school.CNPJ, school.Email, school.Password, school.Street, school.Number, school.ZIP)
	return err
}

func (s *SchoolRepository) ReadSchool(ctx context.Context, cnpj *string) (*types.School, error) {
	sqlQuery := `SELECT id, name, cnpj, email, street, number, zip FROM schools WHERE cnpj = $1 LIMIT 1`
	var school types.School
	err := s.db.QueryRow(sqlQuery, *cnpj).Scan(
		&school.ID,
		&school.Name,
		&school.CNPJ,
		&school.Email,
		&school.Street,
		&school.Number,
		&school.ZIP,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	return &school, nil
}

func (s *SchoolRepository) ReadAllSchools(ctx context.Context) ([]types.School, error) {
	sqlQuery := `SELECT id, name, cnpj, email, street, number, zip FROM schools`

	rows, err := s.db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schools []types.School

	for rows.Next() {
		var school types.School
		err := rows.Scan(&school.ID, &school.Name, &school.CNPJ, &school.Email, &school.Street, &school.Number, &school.ZIP)
		if err != nil {
			return nil, err
		}
		schools = append(schools, school)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return schools, nil
}

func (s *SchoolRepository) UpdateSchool(ctx context.Context) error {
	return nil
}

func (s *SchoolRepository) DeleteSchool(ctx context.Context, cnpj *string) error {
	tx, err := s.db.Begin()
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
	_, err = tx.Exec("DELETE FROM schools WHERE cnpj = $1", cnpj)
	return err
}

func (s *SchoolRepository) AuthSchool(ctx context.Context, school *types.School) (*types.School, error) {
	sqlQuery := `SELECT id, name, cnpj, email, password FROM schools WHERE email = $1 LIMIT 1`
	var schoolData types.School
	err := s.db.QueryRow(sqlQuery, school.Email).Scan(
		&schoolData.ID,
		&schoolData.Name,
		&schoolData.CNPJ,
		&schoolData.Email,
		&schoolData.Password,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	match := schoolData.Password == school.Password
	if !match {
		return nil, fmt.Errorf("email or password wrong")
	}
	schoolData.Password = ""
	return &schoolData, nil
}

func (s *SchoolRepository) CreateInvite(ctx context.Context, invite *types.Invite) error {
	log.Print(invite)
	sqlQuery := `INSERT INTO invites (requester, school, email_school, guest, driver, email_driver, status) VALUES ($1, $2, $3)`
	_, err := s.db.Exec(sqlQuery, invite.School.CNPJ, invite.School.Name, invite.School.Email, invite.Driver.CNH, invite.Driver.Name, invite.Driver.Email, "pending")
	return err
}

func (s *SchoolRepository) DeleteEmployee(ctx context.Context, record_id *int) error {
	tx, err := s.db.Begin()
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
	_, err = tx.Exec("DELETE FROM schools_drivers WHERE record = $1", record_id)
	return err
}

func (s *SchoolRepository) GetEmployees(ctx context.Context, cnpj *string) ([]types.Driver, error) {
	sqlQuery := `SELECT record, name_driver, driver, email_driver FROM schools_drivers WHERE school = $1`

	rows, err := s.db.Query(sqlQuery, *cnpj)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drivers []types.Driver

	for rows.Next() {
		var driver types.Driver
		err := rows.Scan(&driver.ID, &driver.CNH)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return drivers, nil

}

func (s *SchoolRepository) IsEmployee(ctx context.Context, cnh *string) error {

	sqlQuery := `SELECT driver FROM schools_drivers WHERE driver = $1 LIMIT 1`
	var driver types.Driver
	err := s.db.QueryRow(sqlQuery, *cnh).Scan(
		&driver.CNH,
	)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err == sql.ErrNoRows {
		return nil
	}

	return fmt.Errorf("school and driver have a connection")

}
