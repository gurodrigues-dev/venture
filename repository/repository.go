package repository

import (
	"database/sql"
	"fmt"
	"gin/config"
	"gin/models"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func SaveClient(user *models.User, endereco *models.Endereco) (bool, error) {

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

	_, err = db.Exec("INSERT INTO users (id, cpf, rg, name, cnh, email, qrcode, password) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		user.ID, user.CPF, user.RG, user.Name, user.CNH, user.Email, user.URL, user.Password)

	if err != nil {
		return false, err
	}

	_, err = db.Exec("INSERT INTO endereco (rua, cpf, numero, complemento, cidade, estado, cep) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		endereco.Rua, user.CPF, endereco.Numero, endereco.Complemento, endereco.Cidade, endereco.Estado, endereco.CEP)

	return true, nil
}

func FindByCpf(cpf string) (models.GetUser, error) {

	_, err := config.LoadEnvironmentVariables()

	if err != nil {
		return models.GetUser{}, err
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
		return models.GetUser{}, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT u.name, u.email, u.rg, u.cpf, u.cnh, u.qrcode,
		       e.rua, e.numero, e.cep, e.estado, e.cidade, e.complemento
		FROM users u
		INNER JOIN endereco e ON u.cpf = e.cpf
		WHERE u.cpf = $1
	`, cpf)

	if err != nil {
		return models.GetUser{}, err
	}
	defer rows.Close()

	var user models.GetUser
	found := false
	for rows.Next() {
		found = true
		err := rows.Scan(&user.Name, &user.Email, &user.RG, &user.CPF, &user.CNH, &user.URL,
			&user.Endereco.Rua, &user.Endereco.Numero, &user.Endereco.CEP,
			&user.Endereco.Estado, &user.Endereco.Cidade, &user.Endereco.Complemento)
		if err != nil {
			log.Fatal(err)
		}
	}

	if !found {
		return models.GetUser{}, fmt.Errorf("Usuário não encontrado")
	}

	return user, nil

}

func UpdateUser(c *gin.Context, dataToUpdate *models.UpdateUser) (bool, error) {

	query := `
		UPDATE users
		SET email = $1,
		    rua = $2,
		    numero = $3,
		    complemento = $4,
		    cidade = $5,
		    estado = $6,
		    cep = $7
		WHERE cpf = $8
	`

	_, err := db.Exec(query,
		dataToUpdate.Email,
		dataToUpdate.Endereco.Rua,
		dataToUpdate.Endereco.Numero,
		dataToUpdate.Endereco.Complemento,
		dataToUpdate.Endereco.Cidade,
		dataToUpdate.Endereco.Estado,
		dataToUpdate.Endereco.CEP,
	)

	if err != nil {
		return false, err
	}

	return true, nil
}

func DeleteByCpf(cpf string) (bool, error) {

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

	tx, err := db.Begin()
	if err != nil {
		return false, err
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

	_, err = tx.Exec("DELETE FROM endereco WHERE cpf = $1", cpf)
	if err != nil {
		return false, err
	}

	_, err = tx.Exec("DELETE FROM users WHERE cpf = $1", cpf)
	if err != nil {
		return false, err
	}
	return true, nil
}

func VerifyPasswordByCpf(cpf, hash string) (bool, error) {

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

	query := "SELECT password FROM users WHERE cpf = $1"

	err = db.QueryRow(query, cpf).Scan(&storedHashedPassword)
	if err != nil {
		return false, err
	}

	passwordMatch := storedHashedPassword == hash

	return passwordMatch, nil

}

func CheckExistsEmail(email string) (bool, error) {

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
