package network

type Fischannelbinding struct {
	Ifnum     string `json:"ifnum,omitempty"`
	Name      string `json:"name,omitempty"`
	Ownernode int    `json:"ownernode,omitempty"`
}
