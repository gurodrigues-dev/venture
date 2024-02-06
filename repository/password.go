package repository

import (
	"database/sql"
	"fmt"
	"gin/config"
	"gin/models"
)

func ChangePasswordByEmailIdentification(user models.UserResetPassword) (bool, error) {

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

	_, err = db.Exec("UPDATE drivers SET password = $1 WHERE email = $2", user.NewHashPassword, user.Email)

	if err != nil {
		return false, err
	}

	return true, nil

}

func VerifyPasswordByCpf(cpf, table, hash string) (bool, error) {

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

	var storedHashedPassword string

	query := "SELECT password FROM " + table + " WHERE cpf = $1"

	err = db.QueryRow(query, cpf).Scan(&storedHashedPassword)
	if err != nil {
		return false, err
	}

	passwordMatch := storedHashedPassword == hash

	return passwordMatch, nil

}

func VerifyPasswordByCnpj(cnpj, typeSchool, hash *string) (bool, error) {

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

	var storedHashedPassword string

	query := "SELECT password FROM " + *typeSchool + " WHERE cnpj = $1"

	err = db.QueryRow(query, *cnpj).Scan(&storedHashedPassword)
	if err != nil {
		return false, err
	}

	passwordMatch := storedHashedPassword == *hash

	return passwordMatch, nil

}
