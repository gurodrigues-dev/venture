package models

import "github.com/google/uuid"

type CreateDriver struct {
	ID       uuid.UUID
	CPF      string
	RG       string
	Name     string
	CNH      string
	Email    string
	Password string
	URL      string
}

type UpdateDriver struct {
	Email    string
	Endereco struct {
		Rua         string
		Numero      string
		Complemento string
		Cidade      string
		Estado      string
		CEP         string
	}
}

type GetDriver struct {
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
