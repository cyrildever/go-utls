package model_test

import (
	"testing"

	"github.com/cyrildever/go-utls/model"
	"gotest.tools/assert"
)

// TestBits8 ...
func TestBits8(t *testing.T) {
	ref := model.Bits8(0)

	b, _ := model.StringToBits8("00000000")
	assert.Equal(t, ref, b)
}

// TestToBits8 ...
func TestToBits8(t *testing.T) {
	ref := model.Bits8(2)

	bytes1 := []byte{2, 0, 0, 0, 0, 0, 0, 0}
	b1 := model.ToBits8(bytes1)
	assert.Equal(t, ref, b1)

	bytes2 := []byte{2, 100, 0, 0, 0, 0, 0, 0} // No matter what: it will only take the first byte
	b2 := model.ToBits8(bytes2)
	assert.Equal(t, ref, b2)
}

// TestStringToBits8 ...
func TestStringToBits8(t *testing.T) {
	ref := model.Bits8(2)

	b1, _ := model.StringToBits8("00000010")
	assert.Equal(t, ref, b1)

	b2, _ := model.StringToBits8("000010")
	assert.Equal(t, ref, b2)

	_, err := model.StringToBits8("a00010")
	assert.Assert(t, err != nil)
	assert.Equal(t, err.Error(), "not an 8-bit representation")

	_, err = model.StringToBits8("111111111111111111111")
	assert.Assert(t, err != nil)
	assert.Equal(t, err.Error(), "input string too long to be a valid 8-bit representation")
}

// TestBits8Bytes ...
func TestBits8Bytes(t *testing.T) {
	ref := []byte{2}
	b := model.Bits8(2)
	assert.DeepEqual(t, ref, b.Bytes())
}

// TestBits8String ...
func TestBits8String(t *testing.T) {
	ref := "00000010"
	b := model.Bits8(2)
	assert.Equal(t, ref, b.String())
}

// TestBits8ClearAll ...
func TestBits8ClearAll(t *testing.T) {
	b := model.Bits8(1)

	// Before
	assert.Assert(t, b.String() != "00000000")
	assert.Equal(t, b.String(), "00000001")

	b.ClearAll()

	// After
	assert.Equal(t, b.String(), "00000000")
}
