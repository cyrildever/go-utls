package utils_test

import (
	"strings"
	"testing"

	"github.com/cyrildever/go-utls/common/utils"
	"gotest.tools/assert"
)

// TestCapitalize ...
func TestCapitalize(t *testing.T) {
	ref := "My capitalized sentence"
	capitalized := utils.Capitalize("my capitalized sentence")
	assert.Equal(t, capitalized, ref)

	titled := strings.Title("my capitalized sentence")
	assert.Assert(t, capitalized != titled)
	assert.Equal(t, titled, "My Capitalized Sentence")
}

// TestToUTF8 ...
func TestToUTF8(t *testing.T) {
	ref := "just for showing"
	decoded, err := utils.ToUTF8(ref, utils.WINDOWS_1252)
	assert.NilError(t, err)
	assert.Equal(t, ref, decoded)

	_, err = utils.ToUTF8(ref, "wrong format")
	assert.Error(t, err, "unsupported format")
}
