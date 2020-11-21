package memory

import (
	"fmt"

	"github.com/jordan-lumley/a1pos/internal/logger"
	types "github.com/jordan-lumley/a1pos/internal/types"
	"github.com/shirou/gopsutil/mem"
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
			mem, err := mem.VirtualMemory()
			if err != nil {
				logger.Instance().Fatal("FAILED TO GET MEMORY MONITOR")
			}

			memString := fmt.Sprintf("UsedPercent:%f%%", mem.UsedPercent)

			logger.Instance().Info(memString)

			c.Channel <- mem.UsedPercent
		}
	}()
}

// New ...
func New(ch chan float64) (*Monitor, error) {
	return &Monitor{Channel: ch}, nil
}
