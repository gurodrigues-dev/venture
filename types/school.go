package types

type School struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	CNPJ     string `json:"cnpj"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Number   int    `json:"number"`
	CEP      string `json:"cep"`
}
