package config

import (
	"fmt"
	"os"
)

var (
	allMaps map[string]string
)

// Config ...
func initializeConfig() {
	configVar := os.Getenv("A1POS_CONFIG_FILE")
	if configVar == "" {
		fmt.Println("Cannot find the config file path, please make sure A1POS_CONFIG_FILE ENV VAR is set. Exiting...")
		os.Exit(1)
	}

	allMaps = make(map[string]string)

	allMaps["A1POS_CONFIG_FILE"] = configVar
}

// Instance ...
func Instance() map[string]string {
	if allMaps == nil {
		initializeConfig()
	}
	return allMaps
}
