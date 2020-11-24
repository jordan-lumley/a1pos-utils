package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jordan-lumley/a1pos/internal/config"
	"github.com/jordan-lumley/a1pos/pkg/fs"
	"github.com/jordan-lumley/a1pos/pkg/logger"
	"github.com/jordan-lumley/a1pos/pkg/monitor"
)

const (
	refreshPeriod = 1
	refreshScale  = time.Second
)

func main() {
	currentConfig := config.Instance()

	err := fs.EnsureFilePath(currentConfig["A1POS_CONFIG_FILE"])

	file, err := os.OpenFile(currentConfig["A1POS_CONFIG_FILE"],
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger.SetOutput(file)

	cpuMon := monitor.Processor{}
	cpuChan := monitor.Start(cpuMon)

	memMon := monitor.Memory{}
	memChan := monitor.Start(memMon)

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
}
