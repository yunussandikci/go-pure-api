package common

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = logrus.New()
const DefaultLogLevel = logrus.InfoLevel

func init() {
	Logger.SetLevel(getLogLevel())
	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

//getLogLevel Resolved log level from environment variables.
func getLogLevel() logrus.Level {
	logLevel, logLevelPresent := os.LookupEnv("LOG_LEVEL")
	if !logLevelPresent {
		return DefaultLogLevel
	}
	customLogLevel, customLogLevelErr := logrus.ParseLevel(logLevel)
	if customLogLevelErr != nil {
		return DefaultLogLevel
	}
	return customLogLevel
}
