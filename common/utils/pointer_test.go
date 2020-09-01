package utils_test

import (
	"testing"

	"github.com/cyrildever/go-utls/common/utils"
	"gotest.tools/assert"
)

// TestPointer ...
func TestPointer(t *testing.T) {
	val := "This is a value"
	ptr := &val

	assert.Assert(t, utils.IsPointer(ptr))
	assert.Assert(t, utils.IsPointer(val) == false)
	assert.Assert(t, utils.IsValue(val))

	ptrArr := []*string{ptr}
	assert.Assert(t, utils.IsPointer(ptrArr[0]))
}
