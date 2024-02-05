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

func SaveUser(user *models.CreateUser) (bool, error) {

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

	_, err = db.Exec("INSERT INTO users (id, cpf, rg, name, email, password) VALUES ($1, $2, $3, $4, $5, $6)",
		user.ID, user.CPF, user.RG, user.Name, user.Email, user.Password)

	if err != nil {
		return false, err
	}

	_, err = db.Exec("INSERT INTO endereco_users (rua, cpf, numero, complemento, cidade, estado, cep) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		user.Endereco.Rua, user.CPF, user.Endereco.Numero, user.Endereco.Complemento, user.Endereco.Cidade, user.Endereco.Estado, user.Endereco.CEP)

	if err != nil {
		return false, err
	}

	return true, nil
}

func FindUserByCpf(cpf string) (models.GetUser, error) {

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
		SELECT u.name, u.email, u.rg, u.cpf,
		       e.rua, e.numero, e.cep, e.estado, e.cidade, e.complemento
		FROM users u
		INNER JOIN endereco_users e ON u.cpf = e.cpf
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
		err := rows.Scan(&user.Name, &user.Email, &user.RG, &user.CPF,
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

func DeleteUserByCpf(cpf string) (string, error) {

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
	err = tx.QueryRow("SELECT email FROM users WHERE cpf = $1", cpf).Scan(&userEmail)
	if err != nil {
		return "", err
	}

	_, err = tx.Exec("DELETE FROM endereco_users WHERE cpf = $1", cpf)
	if err != nil {
		return "", err
	}

	_, err = tx.Exec("DELETE FROM users WHERE cpf = $1", cpf)
	if err != nil {
		return "", err
	}

	return userEmail, nil
}
