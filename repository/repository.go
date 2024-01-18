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

func SaveDriver(driver *models.CreateDriver, endereco *models.Endereco) (bool, error) {

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

	_, err = db.Exec("INSERT INTO drivers (id, cpf, rg, name, cnh, email, qrcode, password) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		driver.ID, driver.CPF, driver.RG, driver.Name, driver.CNH, driver.Email, driver.URL, driver.Password)

	if err != nil {
		return false, err
	}

	_, err = db.Exec("INSERT INTO endereco (rua, cpf, numero, complemento, cidade, estado, cep) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		endereco.Rua, driver.CPF, endereco.Numero, endereco.Complemento, endereco.Cidade, endereco.Estado, endereco.CEP)

	return true, nil
}

func SaveUser(user *models.CreateUser, endereco *models.Endereco) (bool, error) {

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

	_, err = db.Exec("INSERT INTO endereco (rua, cpf, numero, complemento, cidade, estado, cep) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		endereco.Rua, user.CPF, endereco.Numero, endereco.Complemento, endereco.Cidade, endereco.Estado, endereco.CEP)

	return true, nil
}

func FindDriverByCpf(cpf string) (models.GetDriver, error) {

	_, err := config.LoadEnvironmentVariables()

	if err != nil {
		return models.GetDriver{}, err
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
		return models.GetDriver{}, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT u.name, u.email, u.rg, u.cpf, u.cnh, u.qrcode,
		       e.rua, e.numero, e.cep, e.estado, e.cidade, e.complemento
		FROM drivers u
		INNER JOIN endereco e ON u.cpf = e.cpf
		WHERE u.cpf = $1
	`, cpf)

	if err != nil {
		return models.GetDriver{}, err
	}
	defer rows.Close()

	var user models.GetDriver
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
		return models.GetDriver{}, fmt.Errorf("Usuário não encontrado")
	}

	return user, nil

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

func UpdateDriver(c *gin.Context, dataToUpdate *models.UpdateDriver) (bool, error) {

	query := `
		UPDATE drivers
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

func DeleteByCpf(cpf string) (string, error) {

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

	// Obtenha o e-mail antes de excluir o usuário
	var userEmail string
	err = tx.QueryRow("SELECT email FROM drivers WHERE cpf = $1", cpf).Scan(&userEmail)
	if err != nil {
		return "", err
	}

	_, err = tx.Exec("DELETE FROM endereco WHERE cpf = $1", cpf)
	if err != nil {
		return "", err
	}

	_, err = tx.Exec("DELETE FROM drivers WHERE cpf = $1", cpf)
	if err != nil {
		return "", err
	}

	return userEmail, nil
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

	query := "SELECT password FROM drivers WHERE cpf = $1"

	err = db.QueryRow(query, cpf).Scan(&storedHashedPassword)
	if err != nil {
		return false, err
	}

	passwordMatch := storedHashedPassword == hash

	return passwordMatch, nil

}

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
