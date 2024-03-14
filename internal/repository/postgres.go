package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gin/config"
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

func (p *Postgres) CreateSchool(ctx context.Context) {

}

func (p *Postgres) ReadSchool(ctx context.Context) {

}

func (p *Postgres) UpdateSchool(ctx context.Context) {

}

func (p *Postgres) DeleteSchool(ctx context.Context) {

}

func (p *Postgres) NewPassword(ctx context.Context) {

}
