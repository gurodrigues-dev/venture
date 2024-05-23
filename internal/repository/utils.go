package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gin/config"
	"os"

	_ "github.com/lib/pq"
)

type Repository interface {
	NewPassword(ctx context.Context)
	IsEmailExisting(ctx context.Context, table, email *string) (bool, error)
}

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

func (p *Postgres) IsEmailExisting(ctx context.Context, table, email *string) (bool, error) {
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
