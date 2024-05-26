package types

type Responsible struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CPF        string `json:"cpf"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Street     string `json:"street"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
	ZIP        string `json:"zip"`
}

// childrens of responsible
type Child struct {
	RG          string `json:"rg"`
	Name        string `json:"name"`
	School      string `json:"school"`
	Driver      string `json:"driver"`
	Responsible Responsible
}
