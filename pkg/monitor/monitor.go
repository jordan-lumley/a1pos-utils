package monitor

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// Monitor ...
type Monitor interface {
	start() <-chan float64
}

// Memory ...
type Memory struct{}

// Implement Monitor on Memory
func (m Memory) start() <-chan float64 {
	out := make(chan float64)

	go func() {
		for {
			mem, err := mem.VirtualMemory()
			if err != nil {
			}
			out <- mem.UsedPercent
		}
	}()

	return out
}

// Processor ...
type Processor struct{}

// Implement Monitor on Processer
func (p Processor) start() <-chan float64 {
	out := make(chan float64)

	go func() {
		for {
			cpuPercent, err := cpu.Percent(0, false)
			if err != nil {
			}

			out <- cpuPercent[0]
		}
	}()

	return out
}

// Start ...
func Start(m Monitor) <-chan float64 {
	return m.start()
}
