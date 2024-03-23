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

func (p *Postgres) CreateUser(ctx context.Context) {

}

func (p *Postgres) ReadUser(ctx context.Context) {

}

func (p *Postgres) UpdateUser(ctx context.Context) {

}

func (p *Postgres) DeleteUser(ctx context.Context) {

}

func (p *Postgres) CreateChild(ctx context.Context) {

}

func (p *Postgres) ReadChild(ctx context.Context) {

}

func (p *Postgres) UpdateChild(ctx context.Context) {

}

func (p *Postgres) DeleteChild(ctx context.Context) {

}

func (p *Postgres) CreateDriver(ctx context.Context) {

}

func (p *Postgres) ReadDriver(ctx context.Context) {

}

func (p *Postgres) UpdateDriver(ctx context.Context) {

}

func (p *Postgres) DeleteDriver(ctx context.Context) {

}

func (p *Postgres) CreateSchool(ctx context.Context, school *types.School) error {
	sqlQuery := `INSERT INTO schools (nome, cnpj, email, password, rua, numero, cep) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := p.conn.Exec(sqlQuery, school.Nome, school.CNPJ, school.Email, school.Password, school.Address, school.Number, school.CEP)
	return err
}

func (p *Postgres) ReadSchool(ctx context.Context, id *int) (*types.School, error) {
	sqlQuery := `SELECT name, cnpj, email, password WHERE id = $1 LIMIT 1`
	var school types.School
	err := p.conn.QueryRow(sqlQuery, id).Scan(
		&school.Nome,
		&school.CNPJ,
		&school.Email,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	return &school, nil
}

func (p *Postgres) UpdateSchool(ctx context.Context) error {
	return nil
}

func (p *Postgres) DeleteSchool(ctx context.Context, id *int) error {
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
	_, err = tx.Exec("DELETE FROM schools WHERE id = $1", id)
	return err
}

func (p *Postgres) AuthSchool(ctx context.Context, school *types.School) (*types.School, error) {
	sqlQuery := `SELECT id, nome, cnpj, email, password FROM schools WHERE email = $1 LIMIT 1`
	var schoolData types.School
	err := p.conn.QueryRow(sqlQuery, school.Email).Scan(
		&schoolData.ID,
		&schoolData.Nome,
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

	fmt.Println(err)

	if err != nil {
		return false, err
	}

	return true, nil

}

func (p *Postgres) NewPassword(ctx context.Context) {

}
