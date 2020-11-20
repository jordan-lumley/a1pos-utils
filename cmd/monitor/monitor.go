package monitor

import (
	"fmt"
	"time"

	"github.com/jordan-lumley/a1pos/internal/logger"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

const (
	refreshPeriod = 1
	refreshScale  = time.Second
)

// Execute ...
func Execute() error {
	memChan := make(chan float64)
	cpuChan := make(chan float64)

	go memoryMonitor(memChan)
	go cpuMonitor(cpuChan)

	var memUsage float64
	var cpuUsage float64
	go func() {
		for {
			select {
			case memUsage = <-memChan:
			case cpuUsage = <-cpuChan:
			default:

				fmt.Printf("\rMEMORY USAGE: %f%%/100 CPU USAGE: %f%%/100", memUsage, cpuUsage)
				time.Sleep(refreshScale * refreshPeriod)
			}
		}
	}()

	return nil
}

func memoryMonitor(ch chan float64) {
	for {
		mem, err := mem.VirtualMemory()
		if err != nil {
			logger.Instance().Fatal("FAILED TO GET MEMORY MONITOR")
		}
		memString := fmt.Sprintf("UsedPercent:%f%%", mem.UsedPercent)

		logger.Instance().Info(memString)

		ch <- mem.UsedPercent
	}
}

func cpuMonitor(ch chan float64) {
	for {
		cpuPercent, err := cpu.Percent(0, false)
		if err != nil {
			logger.Instance().Fatal("FAILED TO GET CPU MONITOR")
		}

		cpuString := fmt.Sprintf("%f%%", cpuPercent)

		logger.Instance().Info(cpuString)

		ch <- cpuPercent[0]
	}
}
