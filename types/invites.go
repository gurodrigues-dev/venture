package types

type Invite struct {
	ID        int    `json:"id"`
	Requester string `json:"cnpj"`
	Guest     string `json:"cnh"`
	Status    string `json:"status"`
}
