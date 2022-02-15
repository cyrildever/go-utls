package env_test

import (
	"os"
	"testing"

	"github.com/cyrildever/go-utls/common/env"
	"gotest.tools/assert"
)

// TestEnv ...
func TestEnv(t *testing.T) {
	TEST_ENV := "TEST_ENV"
	os.Setenv(TEST_ENV, "1")

	isSet, boolean := env.GetBool(TEST_ENV)
	assert.Assert(t, isSet)
	assert.Equal(t, boolean, true)

	isSet, integer := env.GetInt(TEST_ENV)
	assert.Assert(t, isSet)
	assert.Equal(t, integer, 1)

	isSet, str := env.GetStr(TEST_ENV)
	assert.Assert(t, isSet)
	assert.Equal(t, str, "1")
}
