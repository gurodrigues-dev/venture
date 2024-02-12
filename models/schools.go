package models

type School struct {
	Name     string `json:"nome"`
	CNPJ     string `json:"cnpj"`
	Rua      string `json:"rua"`
	Numero   string `json:"numero"`
	CEP      string `json:"cep"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginSchool struct {
	CNPJ          string `json:"cnpj"`
	TableOfSearch string `json:"table"`
	Password      string `json:"password"`
}
