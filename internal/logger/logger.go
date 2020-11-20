package logger

import (
	"os"

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

	file, err := os.OpenFile("test.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Info("Failed to log to file, using default stderr")
	} else {
		logger.Out = file
	}

	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	logger.SetReportCaller(true)
}

// Logger ...
func Logger() *logrus.Logger {
	if logger == nil {
		initializeLogger()
	}
	return logger
}
