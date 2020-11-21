package config

import (
	"fmt"
	"os"
)

var (
	// CurrentConfig ...
	allMaps map[string]string
)

// Config ...
func initializeConfig() {
	// configFilePath := os.Getenv("A1POS_CONFIG_FILE")
	// if configFilePath == "" {
	// 	fmt.Println("Cannot find the config file path, please make sure A1POS_CONFIG_FILE ENV VAR is set. Exiting...")
	// 	os.Exit(1)
	// }

	osVar := os.Getenv("OS")
	if osVar == "" {
		fmt.Println("Cannot find the config file path, please make sure OS ENV VAR is set. Exiting...")
		os.Exit(1)
	}

	allMaps = make(map[string]string)

	allMaps["OS"] = "osVar"
}

// Instance ...
func Instance() map[string]string {
	if allMaps == nil {
		initializeConfig()
	}
	return allMaps
}
