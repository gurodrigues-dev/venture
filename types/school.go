package types

type School struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	CNPJ     string `json:"cnpj"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
