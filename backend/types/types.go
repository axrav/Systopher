package types

type Server struct {
	Ip    string `json:"ip_address"`
	Port  string `json:"port"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}
