package models

import "github.com/google/uuid"

type User struct {
	ID    uuid.UUID
	CPF   string
	RG    string
	Name  string
	CNH   string
	Email string
	URL   string
}
