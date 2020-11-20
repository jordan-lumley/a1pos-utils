package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	logger  *log.Logger
	logFile *os.File
	prefix  string
)

type hybridLogger struct{}

func (writer hybridLogger) Write(bytes []byte) (int, error) {
	logFile, err := os.OpenFile("test.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	defer logFile.Close()

	contents := time.Now().UTC().Format("2006-01-02T15:04:05.99Z") + " [DEBUG] " + string(bytes)

	logFile.Write([]byte(contents))

	return fmt.Print(contents)
}

func init() {
	prefix = "A1POS"

	// log.SetOutput(hbLogger)
	logger = log.New(&hybridLogger{}, prefix, log.LstdFlags)
}

// GetLogger ...
func GetLogger() *log.Logger {
	return logger
}
