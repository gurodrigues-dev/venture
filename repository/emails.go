package repository

import (
	"database/sql"
	"fmt"
	"gin/config"
)

func CheckExistsEmailInDrivers(email string) (bool, error) {

	// false -> email not found | true -> email found

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

	var emailDatabase string

	query := "SELECT email FROM drivers WHERE email = $1"

	err = db.QueryRow(query, email).Scan(&emailDatabase)

	if err != nil {
		return false, err
	}

	return true, nil

}

func CheckExistsEmailInUsers(email string) (bool, error) {

	// false -> email not found | true -> email found

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

	var emailDatabase string

	query := "SELECT email FROM users WHERE email = $1"

	err = db.QueryRow(query, email).Scan(&emailDatabase)

	if err != nil {
		return false, err
	}

	return true, nil

}
