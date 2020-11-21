package cpu

import (
	"fmt"

	"github.com/jordan-lumley/a1pos/internal/logger"
	types "github.com/jordan-lumley/a1pos/internal/types"
	"github.com/shirou/gopsutil/cpu"
)

// Monitor ...
type Monitor struct {
	types.IMonitor

	Channel chan float64
}

// Start ...
func (c Monitor) Start() {
	go func() {
		for {
			cpuPercent, err := cpu.Percent(0, false)
			if err != nil {
				logger.Instance().Fatal("FAILED TO GET CPU MONITOR")
			}

			cpuString := fmt.Sprintf("%f%%", cpuPercent)

			logger.Instance().Info(cpuString)

			c.Channel <- cpuPercent[0]
		}
	}()
}

// New ...
func New(ch chan float64) (*Monitor, error) {
	return &Monitor{Channel: ch}, nil
}
