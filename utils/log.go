package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// InitLog init log instant
func InitLog() {
	logger.Out = os.Stdout
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:          false,
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		ForceColors:            true,
		DisableLevelTruncation: true,
	})
	// logger.SetReportCaller(true)
}

// GetLog get log instant
func GetLog() *logrus.Logger {
	return logger
}
