package check

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/docker"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"time"
)

const (
	B  = 1 << 0
	KB = 1 << 10
	MB = 1 << 20
	GB = 1 << 30
)

type HostInfo struct {
	UsageStat
	InfoStat
	CpuInfoStat
	MemStat
	PacketsSent uint64 `json:"packetSent" gorm:"-"` // 上行实时流量
	PacketsRecv uint64 `json:"packetRecv" gorm:"-"` // 下行实时流量
}

//UsageStat 硬盘使用信息
type UsageStat struct {
	DiskTotal       uint64 `json:"diskTotal" gorm:"-"`       // 硬盘总量
	DiskFree        uint64 `json:"diskFree" gorm:"-"`        // 未使用的
	DiskUsed        uint64 `json:"diskUsed" gorm:"-"`        // 使用的
	DiskUsedPercent int    `json:"diskUsedPercent" gorm:"-"` // 已使用百分比
}

//InfoStat 服务操作系统信息
type InfoStat struct {
	Hostname      string `json:"hostname" gorm:"-"`      // 主机名称
	Uptime        uint64 `json:"uptime" gorm:"-"`        // 运行时间
	BootTime      string `json:"bootTime" gorm:"-"`      // 开机时间
	Procs         uint64 `json:"procs" gorm:"-"`         // 进程数量
	OS            string `json:"os" gorm:"-"`            // 操作系统类型
	KernelVersion string `json:"kernelVersion" gorm:"-"` // 操作系统内核版本
	KernelArch    string `json:"kernelArch" gorm:"-"`    // 操作系统架构
	DockerNum     int    `json:"dockerNum" gorm:"-"`     // 运行容器数量
}

// CoreInfoStat cpu核心信息
type CoreInfoStat struct {
	CPU       int32   `json:"cpu" gorm:"-"`       // 编号
	Family    string  `json:"family" gorm:"-"`    // 代数
	Mhz       float64 `json:"mhz" gorm:"-"`       // 主频
	CacheSize int32   `json:"cacheSize" gorm:"-"` // 缓存大小
	Percent   float64 `json:"percent" gorm:"-"`   // 使用率
}

// CpuInfoStat cpu信息
type CpuInfoStat struct {
	Cores  []CoreInfoStat `json:"cores" gorm:"-"` // 核心信息
	Load1  float64        `json:"load1" gorm:"-"`
	Load5  float64        `json:"load5" gorm:"-"`
	Load15 float64        `json:"load15" gorm:"-"`
}

type MemStat struct {
	// Total amount of RAM on this system
	MemTotal uint64 `json:"memTotal" gorm:"-"`

	// RAM available for programs to allocate
	//
	// This value is computed from the kernel specific values.
	MemAvailable uint64 `json:"memAvailable" gorm:"-"`

	// RAM used by programs
	//
	// This value is computed from the kernel specific values.
	MemUsed uint64 `json:"memUsed" gorm:"-"`

	// Percentage of RAM used by programs
	//
	// This value is computed from the kernel specific values.
	MemUsedPercent float64 `json:"memUsedPercent" gorm:"-"`
}

//DiskCheck 服务器硬盘使用量
func DiskCheck() (UsageStat, error) {
	var usage = UsageStat{}
	u, err := disk.Usage("/")
	if err != nil {
		return usage, err
	}

	usage.DiskFree = u.Free
	usage.DiskTotal = u.Total
	usage.DiskUsed = u.Used
	usage.DiskUsedPercent = int(u.UsedPercent)
	return usage, nil
}

//OSCheck 内核检测 操作系统信息获取
func OSCheck() (InfoStat, error) {
	var statInfo = InfoStat{}
	info, err := host.Info()
	if err != nil {
		return statInfo, err
	}

	statInfo.Uptime = info.Uptime / (60 * 60 * 24)
	statInfo.OS = fmt.Sprintf("%s %s %s %s", info.Platform, info.OS, info.PlatformFamily, info.PlatformVersion)
	statInfo.Procs = info.Procs
	statInfo.KernelVersion = info.KernelVersion
	statInfo.KernelArch = info.KernelArch
	statInfo.Hostname = info.Hostname
	statInfo.BootTime = time.Unix(int64(info.BootTime), 0).Format("2006-01-02 15:04:05")
	if statInfo.DockerNum, err = checkDocker(); err != nil {
		return statInfo, err
	}
	return statInfo, nil
}

// CPUCheck cpu使用量
func CPUCheck() (CpuInfoStat, error) {
	var cpuInfo CpuInfoStat
	cpus, err := cpu.Info()
	if err != nil {
		return cpuInfo, err
	}

	cpuInfo.Cores = []CoreInfoStat{}

	pers, err := cpu.Percent(1*time.Second, true)
	if err != nil {
		return cpuInfo, err
	}
	for i, v := range cpus {
		cpuInfo.Cores = append(cpuInfo.Cores, CoreInfoStat{
			CPU:       v.CPU,
			CacheSize: v.CacheSize / 1024,
			Mhz:       v.Mhz / 1000,
			Percent:   pers[i],
			Family:    v.Family,
		})
	}
	a, err := load.Avg()
	if err != nil {
		return cpuInfo, err
	}
	cpuInfo.Load1 = a.Load1
	cpuInfo.Load5 = a.Load5
	cpuInfo.Load15 = a.Load15
	return cpuInfo, nil
}

// RAMCheck 内存使用量
func RAMCheck() (MemStat, error) {
	memStat := MemStat{}
	u, err := mem.VirtualMemory()
	if err != nil {
		return memStat, err
	}

	memStat.MemUsed = u.Used / MB
	memStat.MemUsedPercent = u.UsedPercent
	memStat.MemAvailable = u.Available / MB
	memStat.MemTotal = u.Total / MB
	return memStat, nil
}

func checkDocker() (int, error) {
	ids, err := docker.GetDockerIDList()
	if err != nil {
		return 0, err
	}
	return len(ids), nil
}

func checkIOCounters(info *HostInfo) error {
	IOCounters, err := net.IOCounters(true)
	if err != nil {
		return err
	}
	for _, v := range IOCounters {
		if v.Name == "eth0" {
			info.PacketsSent = v.PacketsSent / MB
			info.PacketsRecv = v.PacketsRecv / MB
		}
	}
	return nil
}

func CheckHost() (HostInfo, error) {
	host := HostInfo{}
	var err error
	if host.MemStat, err = RAMCheck(); err != nil {
		return host, err
	}
	if host.CpuInfoStat, err = CPUCheck(); err != nil {
		return host, err
	}
	if host.InfoStat, err = OSCheck(); err != nil {
		return host, err
	}
	if host.UsageStat, err = DiskCheck(); err != nil {
		return host, err
	}
	if err = checkIOCounters(&host); err != nil {
		return host, err
	}

	return host, nil
}
