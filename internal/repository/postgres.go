package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gin/config"
	"gin/types"
	"os"
)

type Postgres struct {
	conn *sql.DB
}

func NewPostgres() (*Postgres, error) {

	conf := config.Get()

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Password, conf.Database.Name),
	)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	repo := &Postgres{
		conn: db,
	}

	err = repo.migrate(conf.Database.Schema)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (p *Postgres) migrate(filepath string) error {

	schema, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	_, err = p.conn.Exec(string(schema))
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) CreateUser(ctx context.Context, user *types.User) error {
	return nil
}

func (p *Postgres) ReadUser(ctx context.Context, id *int) (*types.User, error) {
	return nil, nil
}

func (p *Postgres) UpdateUser(ctx context.Context) error {
	return nil
}

func (p *Postgres) DeleteUser(ctx context.Context, cpf *string) error {
	return nil
}

func (p *Postgres) AuthUser(ctx context.Context, user *types.User) (*types.User, error) {
	return nil, nil
}

func (p *Postgres) CreateChild(ctx context.Context, child *types.Child, user *types.User) error {
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

func (p *Postgres) CreateDriver(ctx context.Context, driver *types.Driver) error {
	sqlQuery := `INSERT INTO drivers (name, cpf, email, password, cnh, qrcode, street, number, complement, zip) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := p.conn.Exec(sqlQuery, driver.Name, driver.CPF, driver.Email, driver.Password, driver.CNH, driver.QrCode, driver.Street, driver.Number, driver.Complement, driver.ZIP)
	return err
}

func (p *Postgres) ReadDriver(ctx context.Context, cnh *string) (*types.Driver, error) {
	sqlQuery := `SELECT id, name, cpf, cnh, qrcode, email, street, number, zip, complement FROM drivers WHERE cnh = $1 LIMIT 1`
	var driver types.Driver
	err := p.conn.QueryRow(sqlQuery, *cnh).Scan(
		&driver.ID,
		&driver.Name,
		&driver.CPF,
		&driver.CNH,
		&driver.QrCode,
		&driver.Email,
		&driver.Street,
		&driver.Number,
		&driver.ZIP,
		&driver.Complement,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	return &driver, nil
}

func (p *Postgres) UpdateDriver(ctx context.Context) error {
	return nil
}

func (p *Postgres) DeleteDriver(ctx context.Context, cnh *string) error {
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
	_, err = tx.Exec("DELETE FROM drivers WHERE cnh = $1", *cnh)
	return err
}

func (p *Postgres) AuthDriver(ctx context.Context, driver *types.Driver) (*types.Driver, error) {
	sqlQuery := `SELECT id, name, cpf, cnh, email, qrcode, password FROM drivers WHERE email = $1 LIMIT 1`
	var driverData types.Driver
	err := p.conn.QueryRow(sqlQuery, driver.Email).Scan(
		&driverData.ID,
		&driverData.Name,
		&driverData.CPF,
		&driverData.CNH,
		&driverData.Email,
		&driverData.QrCode,
		&driverData.Password,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	match := driverData.Password == driver.Password
	if !match {
		return nil, fmt.Errorf("email or password wrong")
	}
	driverData.Password = ""
	return &driverData, nil
}

func (p *Postgres) CreateSchool(ctx context.Context, school *types.School) error {
	sqlQuery := `INSERT INTO schools (name, cnpj, email, password, street, number, zip) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := p.conn.Exec(sqlQuery, school.Name, school.CNPJ, school.Email, school.Password, school.Street, school.Number, school.ZIP)
	return err
}

func (p *Postgres) ReadSchool(ctx context.Context, cnpj *string) (*types.School, error) {
	sqlQuery := `SELECT id, name, cnpj, email, street, number, zip FROM schools WHERE cnpj = $1 LIMIT 1`
	var school types.School
	err := p.conn.QueryRow(sqlQuery, *cnpj).Scan(
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

func (p *Postgres) ReadAllSchools(ctx context.Context) ([]types.School, error) {
	sqlQuery := `SELECT id, name, cnpj, email, street, number, zip FROM schools`

	rows, err := p.conn.Query(sqlQuery)
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

func (p *Postgres) UpdateSchool(ctx context.Context) error {
	return nil
}

func (p *Postgres) DeleteSchool(ctx context.Context, cnpj *string) error {
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
	_, err = tx.Exec("DELETE FROM schools WHERE cnpj = $1", cnpj)
	return err
}

func (p *Postgres) AuthSchool(ctx context.Context, school *types.School) (*types.School, error) {
	sqlQuery := `SELECT id, name, cnpj, email, password FROM schools WHERE email = $1 LIMIT 1`
	var schoolData types.School
	err := p.conn.QueryRow(sqlQuery, school.Email).Scan(
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

func (p *Postgres) VerifyEmailExists(ctx context.Context, table, email *string) (bool, error) {
	sqlQuery := "SELECT email FROM " + *table + " WHERE email = $1"

	var emailDatabase string

	err := p.conn.QueryRow(sqlQuery, email).Scan(&emailDatabase)

	if err != nil {
		return false, err
	}

	return true, nil

}

func (p *Postgres) NewPassword(ctx context.Context) {

}

func (p *Postgres) CreateInvite(ctx context.Context) {

}

func (p *Postgres) ReadInvite(ctx context.Context) {

}

func (p *Postgres) UpdateInvite(ctx context.Context) {

}

func (p *Postgres) DeleteInvite(ctx context.Context) {

}
