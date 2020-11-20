package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jordan-lumley/a1pos/internal/config"
	"github.com/sirupsen/logrus"
)

const (
	prefix = "A1POS: "
)

var (
	logger *logrus.Logger
)

func initializeLogger() {
	logger = logrus.New()

	initializeLoggerFile()

	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	logger.SetReportCaller(true)
}

func initializeLoggerFile() {
	_, err := os.Stat("logs")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("logs", 0755)
		if errDir != nil {
			panic(err)
		}
	}

	currentConfig, err := config.CurrentConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	logsDir := filepath.Join("logs", filepath.Base(currentConfig.LogConfig.File))
	file, err := os.OpenFile(logsDir,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Info("Failed to log to file, using default stderr")
	} else {
		logger.Out = file
	}
}

// Instance ...
func Instance() *logrus.Logger {
	if logger == nil {
		initializeLogger()
	}
	return logger
}
