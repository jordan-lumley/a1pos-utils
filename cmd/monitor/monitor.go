package monitor

import (
	"fmt"
	"time"

	"github.com/jordan-lumley/a1pos/internal/logger"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

// Execute ...
func Execute() {
	go memoryMonitor()
	go cpuMonitor()
	go diskMonitor()
}

func memoryMonitor() {
	for {
		mem, err := mem.VirtualMemory()
		if err != nil {
			logger.Logger().Fatal("FAILED TO GET MEMORY MONITOR")
		}
		memString := fmt.Sprintf("UsedPercent:%f%%", mem.UsedPercent)

		logger.Logger().Info(memString)
		time.Sleep(time.Second * 5)
	}
}

func cpuMonitor() {
	for {
		cpuPercent, err := cpu.Percent(0, false)
		if err != nil {
			logger.Logger().Fatal("FAILED TO GET CPU MONITOR")
		}

		cpuString := fmt.Sprintf("UsedPercent:%f%%", cpuPercent)

		logger.Logger().Info(cpuString)
		time.Sleep(time.Second * 30)
	}
}

func diskMonitor() {
	for {
		diskCounter, err := disk.IOCounters()
		if err != nil {
			logger.Logger().Fatal("FAILED TO GET DISK MONITOR")
		}

		logger.Logger().Info(diskCounter)
		time.Sleep(time.Second * 30)
	}
}
