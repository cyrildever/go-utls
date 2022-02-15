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

	boolean, wasSet := env.GetBool(TEST_ENV)
	assert.Assert(t, wasSet)
	assert.Equal(t, boolean, true)

	integer, wasSet := env.GetInt(TEST_ENV)
	assert.Assert(t, wasSet)
	assert.Equal(t, integer, 1)

	str, wasSet := env.GetStr(TEST_ENV)
	assert.Assert(t, wasSet)
	assert.Equal(t, str, "1")

	def, wasSet := env.GetStr("DEFAULT_ENV", "my-default-value")
	assert.Assert(t, !wasSet)
	assert.Equal(t, def, "my-default-value")
	assert.Equal(t, os.Getenv("DEFAULT_ENV"), def)
}
