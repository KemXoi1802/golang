package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// InitLog init log instant
func Init(logFileName string, level string) {
	logger.Out = os.Stdout
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:          false,
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		ForceColors:            true,
		DisableLevelTruncation: true,
	})

	
	// f, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	// if err == nil {
	// 	logger.Out = f
	// }

	// if l, err := logrus.ParseLevel(level); err != nil {
	// 	logrus.Fatal("Unknown loglevel ", l)
	// } else {
	// 	logrus.SetLevel(l)
	// }
}

// GetLog get log instant
func GetLog() *logrus.Logger {
	return logger
}