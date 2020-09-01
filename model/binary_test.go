package model_test

import (
	"testing"

	"github.com/cyrildever/go-utls/model"
	"gotest.tools/assert"
)

// TestToBinary ...
func TestToBinary(t *testing.T) {
	bytes := []byte{10, 0, 0, 0, 0, 0, 0, 0}
	ref := model.Binary("1010")
	binary := model.ToBinary(bytes)
	assert.Equal(t, binary, ref)

	zero := model.Binary("00000000")
	zeroBytes := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	zeroBinary := model.ToBinary(zeroBytes)
	assert.Equal(t, zero, zeroBinary)
}

// TestBytes ...
func TestBytes(t *testing.T) {
	op := model.Binary("1010")
	ref := []byte{10, 0, 0, 0, 0, 0, 0, 0}
	assert.DeepEqual(t, op.Bytes(), ref)

	assert.Assert(t, model.Binary("").Bytes() == nil)
}

// TestIsBinaryString ...
func TestIsBinaryString(t *testing.T) {
	str := "100011"
	assert.Assert(t, model.IsBinaryString(str))

	wrongStr := "0abc10"
	assert.Assert(t, model.IsBinaryString(wrongStr) == false)
}
