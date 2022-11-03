package helpers

import (
	"fmt"
	"math"
	"os"
	"os/user"
	"sort"
	"time"

	"github.com/axrav/Systopher/microservice/types"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/showwin/speedtest-go/speedtest"
)

var DataChan = make(chan types.ServerData, 1)

func SortCpu(processes []types.Process) []types.Process {
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].CPU > processes[j].CPU
	})
	return processes
}

func RemoveDuplicates(processes []types.Process) []types.Process {
	keys := make(map[int32]bool)
	list := []types.Process{}
	for _, entry := range processes {
		if _, value := keys[entry.Ppid]; !value {
			keys[entry.Ppid] = true
			list = append(list, entry)
		}
	}
	return list
}
func ErrCheck(error error) {
	if error != nil {
		fmt.Println(error)
	}
}

func TopProcesses() []types.Process {
	user, err := user.Current()
	ErrCheck(err)
	process, err := process.Processes()
	ErrCheck(err)
	var processes []types.Process
	for _, p := range process {
		name, err := p.Name()
		ErrCheck(err)
		memory, err := p.MemoryPercent()
		ErrCheck(err)
		cpu, err := p.CPUPercent()
		ErrCheck(err)
		ppid, err := p.Ppid()
		ErrCheck(err)
		newProcess := types.Process{
			Name:          name,
			MemoryPercent: fmt.Sprintf("%g", (math.Round(float64(memory*100))/100)) + "%",
			CPU:           fmt.Sprintf("%g", math.Round(float64(cpu*100))/100) + "%",
			Pid:           p.Pid,
			Ppid:          ppid,
			User:          user.Username,
		}
		processes = append(processes, newProcess)
	}
	return RemoveDuplicates(SortCpu(processes))[:4]

}

func BasicData() types.ServerData {
	User, _ := user.Current()
	diskdata, _ := disk.Usage("/")
	Total, Used := diskdata.Total, diskdata.Used
	totalProcesses, _ := process.Processes()
	count := len(totalProcesses)
	Cpu, _ := cpu.Info()
	CoreCount := len(Cpu)
	RAM, _ := mem.VirtualMemory()
	TotalRam, UsedRam, TotalSwap, UsedSwap, FreeSwap, MemoryCache := RAM.Total/1024/1024, RAM.Used/1024/1024, RAM.SwapTotal/1024/1024, RAM.SwapCached/1024/1024, RAM.SwapFree/1024/1024, RAM.Cached/1024/1024

	CpuModel := Cpu[0].ModelName
	HostInfo, _ := host.Info()
	Kernel := HostInfo.KernelVersion
	OS := os.Getenv("PRETTY_NAME")
	uptime := HostInfo.Uptime
	days := uptime / (60 * 60 * 24)
	hours := (uptime - (days * 60 * 60 * 24)) / (60 * 60)
	minutes := ((uptime - (days * 60 * 60 * 24)) - (hours * 60 * 60)) / 60
	Uptime_final := fmt.Sprintf("%d days, %d hours, %d minutes", days, hours, minutes)
	Percent, _ := cpu.Percent(time.Second, false)
	CPUUsage := fmt.Sprintf("%g", math.Round(float64(Percent[0]*100))/100) + "%"
	TopProcesses := TopProcesses()

	return types.ServerData{
		Ip:             HostInfo.Hostname,
		OS:             OS,
		Kernel:         Kernel,
		Uptime:         Uptime_final,
		CPUUsage:       CPUUsage,
		CPU:            CpuModel,
		Core:           fmt.Sprintf("%d", CoreCount),
		TotalMemory:    fmt.Sprintf("%d Mb", TotalRam),
		UsedMemory:     fmt.Sprintf("%d Mb", UsedRam),
		FreeMemory:     fmt.Sprintf("%d Mb", TotalRam-UsedRam),
		TotalSwapiness: fmt.Sprintf("%d Mb", TotalSwap),
		SwapUsed:       fmt.Sprintf("%d Mb", UsedSwap),
		FreeSwap:       fmt.Sprintf("%d Mb", FreeSwap),
		Cache:          fmt.Sprintf("%d Mb", MemoryCache),
		Disk:           fmt.Sprintf("%g Gb", math.Round(float64((Total*100))/1024/1024/1024)/100),
		DiskUsed:       fmt.Sprintf("%g Gb", math.Round(float64((Used*100))/1024/1024/1024)/100),
		Processes:      fmt.Sprintf("%d", count),
		TopProcesses:   TopProcesses,
		User:           User.Username,
	}

}

func SendServerData() types.ServerData {
	user, _ := speedtest.FetchUserInfo()

	serverList, _ := speedtest.FetchServers(user)
	targets, _ := serverList.FindServer([]int{})
	init_data := BasicData()
	for _, target := range targets {
		target.PingTest()
		target.DownloadTest(false)
		target.UploadTest(false)

		init_data.DownloadSpeed = fmt.Sprintf("%g Mbps", math.Round(float64(target.DLSpeed*100))/100)
		init_data.UploadSpeed = fmt.Sprintf("%g Mbps", math.Round(float64(target.ULSpeed*100))/100)
		init_data.Ping = target.Latency.String()

	}
	return init_data

}
