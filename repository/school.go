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

func DeleteSchoolByCnpj(cpf *string) error {

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

	return nil

}
