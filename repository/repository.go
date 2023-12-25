package repository

import (
	"database/sql"
	"fmt"
	"gin/config"
	"gin/models"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InsertNewUser(user models.User, endereco models.Endereco) (bool, error) {

	_, err := config.LoadEnvironmentVariables()

	if err != nil {
		return false, err
	}

	var (
		userdb   = config.GetUserDatabase()
		port     = config.GetPortDatabase()
		host     = config.GetHostDatabase()
		password = config.GetPasswordDatabase()
		dbname   = config.GetNameDatabase()
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, userdb, password, dbname)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return false, err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users (id, cpf, rg, name, cnh, email, qrcode) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		user.ID, user.CPF, user.RG, user.Name, user.CNH, user.Email, user.URL)

	if err != nil {
		return false, err
	}

	_, err = db.Exec("INSERT INTO endereco (rua, cpf, numero, complemento, cidade, estado, cep) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		endereco.Rua, user.CPF, endereco.Numero, endereco.Complemento, endereco.Cidade, endereco.Estado, endereco.CEP)

	return true, nil
}

func GetUser(id int64) (bool, error) {
	return true, nil
}

func GetAllUsers() (bool, error) {
	return true, nil
}

func UpdateUser(id int64) (bool, error) {
	return true, nil
}

func DeleteUser(id int64) (bool, error) {
	return true, nil
}
