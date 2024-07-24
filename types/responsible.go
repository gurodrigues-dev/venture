package types

type Responsible struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	CPF           string `json:"cpf"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	Street        string `json:"street"`
	Number        string `json:"number"`
	Complement    string `json:"complement"`
	ZIP           string `json:"zip"`
	Token         string `json:"cardtoken"`
	CustomerID    string `json:"customerId"`
	PaymentMethod string `json:"paymentMethod"`
}

// childrens of responsible
type Child struct {
	RG          string `json:"rg"`
	Name        string `json:"child"`
	School      string `json:"school"`
	Driver      string `json:"driver"`
	Shift       string `json:"shift"`
	Responsible Responsible
}
