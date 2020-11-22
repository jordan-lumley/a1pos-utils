package monitor

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

// Monitor ...
type Monitor interface {
	start()
}

// Memory ...
type Memory struct {
	Output chan float64
}

// Implement Monitor on Memory
func (m Memory) start() {
	for {
		mem, err := mem.VirtualMemory()
		if err != nil {
		}

		m.Output <- mem.UsedPercent
	}
}

// Processor ...
type Processor struct {
	Output chan float64
}

// Implement Monitor on Processer
func (p Processor) start() {
	for {
		cpuPercent, err := cpu.Percent(0, false)
		if err != nil {
		}

		p.Output <- cpuPercent[0]
	}
}

// Start ...
func Start(m Monitor) {
	m.start()
}
