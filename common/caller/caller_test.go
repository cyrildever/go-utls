package caller_test

import (
	"os"
	"testing"

	"github.com/cyrildever/go-utls/common/caller"
	"gotest.tools/assert"
)

// TestCaller ...
func TestCaller(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	calling := func() {
		file, line, ok := caller.Get()
		if !ok {
			t.Fatal("impossible to use the caller")
		}
		assert.Equal(t, file, dir+"/caller_test.go")
		assert.Equal(t, line, 29)

		str := caller.AsString()
		assert.Assert(t, str != "")
		assert.Equal(t, str, dir+"/caller_test.go#29")
	}
	calling() // <- This is where it's actually called
}
