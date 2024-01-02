package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	CPF      string
	RG       string
	Name     string
	CNH      string
	Email    string
	Password string
	URL      string
}

type UpdateUser struct {
	Email    string
	Password string
	Endereco struct {
		Rua         string
		Numero      string
		Complemento string
		Cidade      string
		Estado      string
		CEP         string
	}
}

type GetUser struct {
	CPF      string
	RG       string
	Name     string
	CNH      string
	Email    string
	URL      string
	Endereco struct {
		Rua         string
		Numero      string
		Complemento string
		Cidade      string
		Estado      string
		CEP         string
	}
}

type AuthUser struct {
	CPF string
}
