package service

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"gin/config"
	"gin/internal/repository"
	"gin/types"

	_ "github.com/lib/pq"
)

func setupTestDB(t *testing.T) (*sql.DB, *SchoolService) {
	t.Helper()

	config, err := config.Load("../../config/config.yaml")
	if err != nil {
		t.Fatalf("falha ao carregar a configuração: %v", err)
	}

	db, err := sql.Open("postgres", newPostgres(config.Database))
	if err != nil {
		t.Fatalf("falha ao conectar ao banco de dados: %v", err)
	}

	schoolRepository := repository.NewSchoolRepository(db)
	schoolService := NewSchoolService(schoolRepository)

	return db, schoolService
}

func mockSchool() *types.School {
	return &types.School{
		Name:       "EE Professor Armando Gomes de Araujo",
		CNPJ:       "48480362000153",
		Email:      "gustavorodrigueslima2004@gmail.com",
		Password:   "123teste",
		Street:     "Rua Alfredo ",
		Number:     "1",
		Complement: "",
		ZIP:        "08110220",
	}
}

// func deleteMockSchool(t *testing.T, db *sql.DB, cnpj string) {
// 	t.Helper()
// 	query := "DELETE FROM schools WHERE cnpj = $1"
// 	_, err := db.Exec(query, cnpj)
// 	if err != nil {
// 		t.Fatalf("falha ao deletar mock da escola: %v", err)
// 	}
// }

func TestCreateSchool(t *testing.T) {
	db, schoolService := setupTestDB(t)
	defer db.Close()

	schoolMock := mockSchool()
	// deleteMockSchool(t, db, schoolMock.CNPJ)

	err := schoolService.CreateSchool(context.Background(), schoolMock)
	if err != nil {
		t.Errorf("Erro ao criar escola: %v", err)
	}
}

func TestGetSchool(t *testing.T) {
	db, schoolService := setupTestDB(t)
	defer db.Close()

	schoolMock := mockSchool()

	// schoolData is the school struct returned from the database
	schoolData, err := schoolService.ReadSchool(context.Background(), &schoolMock.CNPJ)

	if err != nil {
		t.Errorf("Erro ao fazer leitura da escola: %v", err.Error())
	}

	// transforming the mock data that is empty or will be returned as empty,
	// the same as that returned from the database so that the validation is done faithfully
	schoolMock.Password = ""
	schoolMock.ID = schoolData.ID

	if *schoolMock != *schoolData {
		t.Error("Mock é diferente do user retornado do banco")
	}
}

func TestGetAllSchools(t *testing.T) {

	db, schoolService := setupTestDB(t)
	defer db.Close()

	schools, err := schoolService.ReadAllSchools(context.Background())

	if err != nil {
		t.Errorf("Erro ao encontrar lista das escolas: %v", err.Error())
	}

	if reflect.TypeOf(schools) != reflect.TypeOf([]types.School{}) {
		t.Errorf("Não foi retornado uma lista de escolas: %v", err.Error())
	}

}

func TestUpdateSchool(t *testing.T) {

	db, schoolService := setupTestDB(t)
	defer db.Close()

	newSchool := types.School{
		Name:       "EE Professor Armando Gomes Araujo",
		CNPJ:       "48480362000153",
		Email:      "gustavorodrigueslima2004@gmail.com",
		Password:   "13867443",
		Street:     "Rua Alfredo Pariense",
		Number:     "1",
		Complement: "",
		ZIP:        "08110220",
	}

	err := schoolService.UpdateSchool(context.Background(), &newSchool)
	if err != nil {
		t.Errorf("Erro ao atualizar escola: %v", err)
	}

}

func TestAuthSchool(t *testing.T) {

	db, schoolService := setupTestDB(t)
	defer db.Close()

	schoolMock := mockSchool()

	schoolData, err := schoolService.AuthSchool(context.Background(), schoolMock)

	if err != nil {
		t.Errorf("Erro ao fazer login da escola: %v", err.Error())
	}

	if reflect.TypeOf(schoolData) != reflect.TypeOf(schoolMock) {
		t.Errorf("Não foi retornado os dados de login da escola: %v", err.Error())
	}

}

func TestDeleteSchool(t *testing.T) {

	db, schoolService := setupTestDB(t)
	defer db.Close()

	schoolMock := mockSchool()

	err := schoolService.DeleteSchool(context.Background(), &schoolMock.CNPJ)

	if err != nil {
		t.Errorf("Erro ao deletar escola: %v", err.Error())
	}

}
