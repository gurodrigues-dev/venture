package service

import (
	"context"
	"database/sql"
	"gin/config"
	"gin/internal/repository"
	"gin/types"
	"testing"

	_ "github.com/lib/pq"
)

func newPostgres(dbConfig config.Database) string {
	return "user=" + dbConfig.User +
		" password=" + dbConfig.Password +
		" dbname=" + dbConfig.Name +
		" host=" + dbConfig.Host +
		" port=" + dbConfig.Port +
		" sslmode=disable"
}

func mockResponsible() *types.Responsible {
	return &types.Responsible{
		Name:       "Alvares de Lima",
		CPF:        "22321279826",
		Email:      "testekauabarbosa@kaubarbosa.dev",
		Password:   "123teste",
		Street:     "Rua Parecue",
		Number:     "55",
		Complement: "Sobrado",
		ZIP:        "02636070",
	}
}

func deleteMockResponsible(db *sql.DB, cpf string) error {
	query := "DELETE FROM responsibles WHERE cpf = $1"
	_, err := db.Exec(query, cpf)
	return err
}

func TestCreateResponsible(t *testing.T) {
	config, err := config.Load("../../config/config.yaml")
	if err != nil {
		t.Fatalf("falha ao carregar a configuração: %v", err)
	}

	db, err := sql.Open("postgres", newPostgres(config.Database))
	if err != nil {
		t.Fatalf("falha ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	responsibleRepository := repository.NewResponsibleRepository(db)
	responsibleService := NewResponsibleService(responsibleRepository)

	responsible := mockResponsible()

	err = deleteMockResponsible(db, responsible.CPF)

	if err != nil {
		t.Errorf("responsible mock doesn't deleted")
	}

	err = responsibleService.CreateResponsible(context.Background(), responsible)

	if err != nil {
		t.Errorf("Erro ao criar responsável: %v", err)
	}
}

func TestReadResponsible(t *testing.T) {

}

func TestUpdateResponsible(t *testing.T) {

}

func TestDeleteResponsible(t *testing.T) {

}

func TestAuthResponsible(t *testing.T) {

}

func TestCreateChild(t *testing.T) {

}

func TestReadChildren(t *testing.T) {

}

func TestUpdateChild(t *testing.T) {

}

func TestDeleteChild(t *testing.T) {

}

func TestIsSponsor(t *testing.T) {

}

func TestParserJwtResponsible(t *testing.T) {

}

func TestCreateTokenJWTResponsible(t *testing.T) {

}
