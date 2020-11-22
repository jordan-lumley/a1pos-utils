package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

func initializeLogger() {
	logger = logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	logger.SetReportCaller(true)
}

// SetOutput ...
func SetOutput(w io.Writer) {
	Instance().Out = w
}

// Instance ...
func Instance() *logrus.Logger {
	if logger == nil {
		initializeLogger()
	}
	return logger
}
