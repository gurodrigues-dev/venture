package types

type Invite struct {
	ID        int    `json:"id"`
	Requester string `json:"requester"`
	Guest     string `json:"guest"`
	Status    string `json:"status"`
}
