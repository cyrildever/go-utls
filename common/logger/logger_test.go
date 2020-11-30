package logger_test

import (
	"testing"

	"github.com/cyrildever/go-utls/common/logger"
)

func TestLog(t *testing.T) {
	log := logger.Init("logger_test", "TestLogger", "test.log")
	log.Debug("Test debug")
	log.Info("Test info")
	log.Warn("Test warn")
	log.Error("Test error")
	log.Crit("Test crit")

	// assert.Assert(t, false)
}
