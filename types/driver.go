package types

type Driver struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	CPF        string `json:"cpf"`
	CNH        string `json:"cnh"`
	QrCode     string `json:"qrcode"`
	Street     string `json:"street"`
	Number     int    `json:"number"`
	ZIP        string `json:"zip"`
	Complement string `json:"complement"`
}
