package types

type Child struct {
	RG          string `json:"rg"`
	Name        string `json:"name"`
	School      string `json:"school"`
	Driver      string `json:"driver"`
	Responsible Responsible
}
