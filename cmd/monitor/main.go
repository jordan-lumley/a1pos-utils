package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jordan-lumley/a1pos/internal/config"
	"github.com/jordan-lumley/a1pos/internal/logger"
	cpuInternal "github.com/jordan-lumley/a1pos/internal/monitors/cpu"
	"github.com/jordan-lumley/a1pos/internal/monitors/memory"
	"github.com/jordan-lumley/a1pos/internal/service"
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

	cpuMon, err := cpuInternal.New(cpuChan)
	if err != nil {
		panic(err)
	}
	go cpuMon.Start()

	memMon, err := memory.New(memChan)
	if err != nil {
		panic(err)
	}
	go memMon.Start()

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
