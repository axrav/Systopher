package types

type Server struct {
	Ip    string `json:"ip_address"`
	Port  string `json:"port"`
	Name  string `json:"name" `
	Owner string `json:"owner" `
	Token string `json:"token" `
}
type ServerData struct {
	Ip             string    `json:"ip_address"`
	TotalMemory    string    `json:"total_memory"`
	Ping           string    `json:"ping"`
	UsedMemory     string    `json:"used_memory"`
	FreeMemory     string    `json:"free_memory"`
	CPU            string    `json:"cpu"`
	Core           string    `json:"core"`
	Disk           string    `json:"disk"`
	DiskUsed       string    `json:"disk_used"`
	DownloadSpeed  string    `json:"network_download"`
	UploadSpeed    string    `json:"network_upload"`
	Uptime         string    `json:"uptime"`
	Kernel         string    `json:"kernel"`
	OS             string    `json:"os"`
	TotalSwapiness string    `json:"swapiness"`
	SwapUsed       string    `json:"swap_used"`
	FreeSwap       string    `json:"swap_free"`
	Processes      string    `json:"processes"`
	CPUUsage       string    `json:"cpu_usage"`
	TopProcesses   []Process `json:"top_processes"`
	User           string    `json:"user"`
	Cache          string    `json:"cache"`
}

type Process struct {
	Name          string `json:"name"`
	CPU           string `json:"cpu"`
	MemoryPercent string `json:"mem"`
	Pid           int32  `json:"pid"`
	User          string `json:"owner"`
	Ppid          int32  `json:"ppid"`
}
type OTPResponse struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UniqueID string `json:",omitempty"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

type UserData struct {
	Email    string   `json:"email"`
	Username string   `json:"username"`
	UniqueID string   `json:"unique_id"`
	Servers  []Server `json:"servers"`
}

type Email struct {
	Email string `json:"email"`
}

type Password struct {
	Password    string `json:"password,omitempty"`
	NewPassword string `json:"new_password"`
	OTP         string `json:"otp,omitempty"`
}
