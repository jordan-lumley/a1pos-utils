package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	// CurrentConfig ...
	currentConfig Object

	configFileContents []byte
)

// Object ...
type Object struct {
	LogConfig logConfig `json:"log"`
}

// LogConfig ...
type logConfig struct {
	File string `json:"file"`
}

// Config ...
func Config() error {
	configFilePath, err := filepath.Abs("./config.json")
	if err != nil {
		return err
	}

	configFile, err := os.Open(configFilePath)
	if err != nil {
		return err
	}

	c, err := ioutil.ReadAll(configFile)
	if err != nil {
		return err
	}

	configFileContents = c

	json.Unmarshal(configFileContents, &currentConfig)

	return nil
}

// CurrentConfig ...
func CurrentConfig() (Object, error) {
	return currentConfig, nil
}
