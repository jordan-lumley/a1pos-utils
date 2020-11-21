package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jordan-lumley/a1pos/internal/config"
	"github.com/jordan-lumley/a1pos/internal/logger"
	"github.com/jordan-lumley/a1pos/internal/service"
	"github.com/jordan-lumley/a1pos/pkg/monitor/memory"
	processor "github.com/jordan-lumley/a1pos/pkg/monitor/processor"
	"github.com/jordan-lumley/a1pos/pkg/monitor/types"
)

const (
	refreshPeriod = 1
	refreshScale  = time.Second
)

func main() {
	_, err := os.Stat("logs")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("logs", 0755)
		if errDir != nil {
			fmt.Println(err)
			panic(err)
		}
	}

	currentConfig := config.Instance()
	fmt.Println(currentConfig)

	logsDir := filepath.Join("logs", filepath.Base("currentConfig.LogConfig.File"))
	file, err := os.OpenFile(logsDir,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	logger.SetOutput(file)

	memChan := make(chan float64)
	cpuChan := make(chan float64)

	go Start(processor.Processor{Channel: cpuChan})

	go Start(memory.Memory{Channel: memChan})

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

	service.Run()
}

// Start ...
func Start(monitor types.IMonitor) {
	monitor.Start()
}
