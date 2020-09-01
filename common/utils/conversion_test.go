package utils_test

import (
	"testing"

	"github.com/cyrildever/go-utls/common/utils"
	"gotest.tools/assert"
)

//--- These variables are equivalent to each other
var integer uint64 = 180969
var bytearray = []byte{233, 194, 2, 0, 0, 0, 0, 0}

// TestUintToByteArray ...
func TestUintToByteArray(t *testing.T) {
	enc := utils.UintToByteArray(integer)
	assert.DeepEqual(t, bytearray, enc)
}

// TestByteArrayToUint ...
func TestByteArrayToUint(t *testing.T) {
	dec := utils.ByteArrayToUint(bytearray)
	assert.Assert(t, dec == integer)
}

//--- These variables are also equivalent to one another
var u8 uint8 = 255
var b8 = []byte{255}

// TestUint8ToByteArray ...
func TestUint8ToByteArray(t *testing.T) {
	enc := utils.Uint8ToByteArray(u8)
	assert.DeepEqual(t, b8, enc)
}

// TestByteArrayToUint8 ...
func TestByteArrayToUint8(t *testing.T) {
	dec := utils.ByteArrayToUint8(b8)
	assert.Assert(t, dec == u8)
}

// TestIntToByteArray ...
func TestIntToByteArray(t *testing.T) {
	i := -1
	b := utils.IntToByteArray(i)
	assert.DeepEqual(t, b, []byte{255, 255, 255, 255, 255, 255, 255, 255})
	assert.DeepEqual(t, i, utils.ByteArrayToInt(b))
}
