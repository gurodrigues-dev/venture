package types

type ValidaInvite struct {
	Requester string `json:"requester"`
	Guest     string `json:"guest"`
}

type Invite struct {
	ID     int    `json:"id"`
	School School `json:"school"`
	Driver Driver `json:"driver"`
	Status string `json:"status"`
}

type Sponsor struct {
	ID     int    `json:"id"`
	Driver Driver `json:"driver"`
	Child  Child  `json:"child"`
	School School `json:"school"`
	Status string `json:"status"`
}
