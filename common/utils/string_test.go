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

// TestReverse ...
func TestReverse(t *testing.T) {
	ref := "desrever"
	reversed := utils.Reverse("reversed")
	assert.Equal(t, reversed, ref)
}
