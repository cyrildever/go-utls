package model_test

import (
	"testing"

	"github.com/cyrildever/go-utls/model"
	"gotest.tools/assert"
)

// TestIsHexString ...
func TestIsHexString(t *testing.T) {
	ref := "1234"
	assert.Assert(t, model.IsHexString(ref))

	short := "123"
	assert.Assert(t, !model.IsHexString(short))

	wrong := "wrong hex string"
	assert.Assert(t, !model.IsHexString(wrong))
}
