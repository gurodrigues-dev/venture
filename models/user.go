package models

import "github.com/google/uuid"

type CreateUser struct {
	ID       uuid.UUID
	CPF      string
	RG       string
	Name     string
	Email    string
	Password string
}

type GetUser struct {
	CPF      string
	RG       string
	Name     string
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

type InfoUserToDriver struct {
	URL  string
	Info GetUser
}

type UpdateUser struct {
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

type AuthUser struct {
	CPF string
}

type UserInfoToResetPassword struct {
	Token string
	Email string
}

type UserResetPassword struct {
	Email           string
	NewHashPassword string
}
