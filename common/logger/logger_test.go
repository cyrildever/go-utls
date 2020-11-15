package logger_test

import (
	"testing"

	"github.com/cyrildever/go-utls/common/logger"
)

// TestLogger ...
func TestLogger(t *testing.T) {
	log := logger.Init("logger_test", "TestLogger", "test.log")
	log.Debug("Test debug")
	log.Info("Test info")
	log.Error("Test error")
	log.Crit("Test crit")

	// assert.Assert(t, false)
}
