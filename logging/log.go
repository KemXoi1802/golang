package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// Init init log instant
func Init(logFileName string, level string) {
	logger.Out = os.Stdout
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:          false,
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		ForceColors:            true,
		DisableLevelTruncation: true,
	})

	// logger.SetReportCaller(true)

	// f, err := os.OpenFile(logFileName, os.O_CREATE, 0755)
	f, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err == nil {
		logger.SetOutput(f)
	}

	logger.SetLevel(logrus.DebugLevel)

}

// GetLog get log instant
func GetLog() *logrus.Logger {
	return logger
}
