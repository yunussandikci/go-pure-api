package common

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLogger_GetLogLevelWhenProvidedByEnvironment(t *testing.T) {
	//given
	os.Setenv("LOG_LEVEL", "trace")
	defer os.Unsetenv("LOG_LEVEL")

	//when
	level := getLogLevel()

	//then
	assert.Equal(t, logrus.TraceLevel, level)
}

func TestLogger_GetLogLevelWhenDefault(t *testing.T) {
	//given
	defer os.Unsetenv("LOG_LEVEL")

	//when
	level := getLogLevel()

	//then
	assert.Equal(t, logrus.InfoLevel, level)
}