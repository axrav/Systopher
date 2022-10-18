package types

type Server struct {
	Ip    string `json:"ip_address"`
	Port  string `json:"port"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
	Token string `json:"token"`
}
type ServerData struct {
	Ip              string `json:"ip_address"`
	TotalMemory     string `json:"total_memory"`
	Ping            string `json:"ping"`
	UsedMemory      string `json:"ram"`
	FreeMemory      string `json:"free_memory"`
	CPU             string `json:"cpu"`
	Core            string `json:"core"`
	Disk            string `json:"disk"`
	DiskUsed        string `json:"disk_used"`
	NetworkDownload string `json:"network_download"`
	NetworkUpload   string `json:"network_upload"`
	Uptime          string `json:"uptime"`
	Kernel          string `json:"kernel"`
	OS              string `json:"os"`
	TotalSwapiness  string `json:"swapiness"`
	SwapUsed        string `json:"swap_used"`
	FreeSwap        string `json:"swap_free"`
	Processes       string `json:"processes"`
	CPUUsage        string `json:"cpu_usage"`
	TopProcesses    string `json:"top_processes"`
	User            string `json:"user"`
	Cache           string `json:"cache"`
}

type OTPResponse struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
