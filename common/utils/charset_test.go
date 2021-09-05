package utils_test

import (
	"testing"

	"github.com/cyrildever/go-utls/common/utils"
	"gotest.tools/assert"
)

// TestToUTF8 ...
func TestToUTF8(t *testing.T) {
	ref := "just for showing"
	decoded, err := utils.ToUTF8(ref, utils.WINDOWS_1252)
	assert.NilError(t, err)
	assert.Equal(t, ref, decoded)

	_, err = utils.ToUTF8(ref, "wrong format")
	assert.Error(t, err, "unsupported encoding")

	decoded, err = utils.ToUTF8(ref, "UTF-8")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, decoded, ref)
}
