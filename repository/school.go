package repository

import (
	"database/sql"
	"fmt"
	"gin/config"
	"gin/models"
)

func SaveSchool(school *models.School) error {

	_, err := config.LoadEnvironmentVariables()

	if err != nil {
		return err
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
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO schools (nome, cnpj, email, password, rua, numero, cep) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		school.Name, school.CNPJ, school.Email, school.Password, school.Rua, school.Numero, school.CEP)

	if err != nil {
		return err
	}

	return nil

}

func FindSchoolByName(name *string) {

	return

}

func UpdateSchool() {

	return

}

func DeleteSchoolByCnpj(cnpj *string) (string, error) {

	_, err := config.LoadEnvironmentVariables()

	if err != nil {
		return "", err
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
		return "", err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return "", err
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

	var userEmail string
	err = tx.QueryRow("SELECT email FROM schools WHERE cpf = $1", *cnpj).Scan(&userEmail)
	if err != nil {
		return "", err
	}

	_, err = tx.Exec("DELETE FROM schools WHERE cpf = $1", *cnpj)
	if err != nil {
		return "", err
	}

	return userEmail, nil

}
