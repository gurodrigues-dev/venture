package models

type School struct {
	Name   string `json:"name"`
	CNPJ   string `json:"cnpj"`
	Rua    string `json:"rua"`
	Numero string `json:"numero"`
	CEP    string `json:"cep"`
	Email  string `json:"email"`
}
