package types

type ServerData struct {
	Ip              string    `json:"ip_address"`
	TotalMemory     string    `json:"total_memory"`
	Ping            string    `json:"ping"`
	UsedMemory      string    `json:"used_memory"`
	FreeMemory      string    `json:"free_memory"`
	CPU             string    `json:"cpu"`
	Core            string    `json:"core"`
	Disk            string    `json:"disk"`
	DiskUsed        string    `json:"disk_used"`
	NetworkDownload string    `json:"network_download"`
	NetworkUpload   string    `json:"network_upload"`
	Uptime          string    `json:"uptime"`
	Kernel          string    `json:"kernel"`
	OS              string    `json:"os"`
	TotalSwapiness  string    `json:"swapiness"`
	SwapUsed        string    `json:"swap_used"`
	FreeSwap        string    `json:"swap_free"`
	Processes       string    `json:"processes"`
	CPUUsage        string    `json:"cpu_usage"`
	TopProcesses    []Process `json:"top_processes"`
	User            string    `json:"user"`
	Cache           string    `json:"cache"`
}

type Process struct {
	Name          string `json:"name"`
	CPU           string `json:"cpu"`
	MemoryPercent string `json:"mem"`
	Pid           int32  `json:"pid"`
	User          string `json:"owner"`
	Ppid          int32  `json:"ppid"`
}
