package logger_test

import (
	"testing"

	"github.com/cyrildever/go-utls/common/logger"
	"gotest.tools/assert"
)

// TestLog ...
func TestLog(t *testing.T) {
	log := logger.Init("logger_test", "TestLogger", "./test.log")
	log.Debug("Test debug")
	log.Info("Test info")
	log.Warn("Test warn")
	log.Error("Test error")
	log.Crit("Test crit")

	// assert.Assert(t, false)
}

// TestLoggerTimeFromLine ...
func TestLoggerGetTimeFromLine(t *testing.T) {
	logLine := `2021/06/16 08:30:37.220607 DEBUG msg={Test debug}	service=logger_test	context=TestLogger	[logger_test.go:11]`
	found := logger.GetTimeFromLine(logLine)
	assert.Equal(t, found.Year(), 2021)
	assert.Equal(t, found.Month().String(), "June")
	assert.Equal(t, found.Day(), 16)
	assert.Equal(t, found.Hour(), 8)
	assert.Equal(t, found.Minute(), 30)
	assert.Equal(t, found.Second(), 37)
}
