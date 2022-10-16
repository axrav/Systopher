package types

type Server struct {
	Ip    string `json:"ip_address"`
	Port  string `json:"port"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}
type ServerData struct {
	TotalMemory     string `json:"total_memory"`
	InternetSpeed   string `json:"internet_speed"`
	UsedMemory      string `json:"ram"`
	CPU             string `json:"cpu"`
	Core            string `json:"core"`
	Disk            string `json:"disk"`
	NetworkDownload string `json:"network_download"`
	NetworkUpload   string `json:"network_upload"`
	Uptime          string `json:"uptime"`
	Kernel          string `json:"kernel"`
	OS              string `json:"os"`
	Swapiness       string `json:"swapiness"`
	Processes       string `json:"processes"`
	CPUUsage        string `json:"cpu_usage"`
	CPUGovernor     string `json:"cpu_governor"`
	TopProcesses    string `json:"top_processes"`
	User            string `json:"user"`
	Cache           string `json:"cache"`
}

type OTPResponse struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}
