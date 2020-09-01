package xor_test

import (
	"testing"

	"github.com/cyrildever/go-utls/common/xor"
	"gotest.tools/assert"
)

// TestXor ...
func TestXor(t *testing.T) {
	xor1, _ := xor.Bytes([]byte{4}, []byte{5})
	assert.DeepEqual(t, []byte{1}, xor1)

	xor2, _ := xor.String("a", "b")
	assert.Equal(t, "", xor2)
	assert.Equal(t, []rune(xor2)[0], rune(3))

	xor3, _ := xor.String("ab", "cd")
	assert.Equal(t, "", xor3)
	assert.DeepEqual(t, []rune(xor3), []rune{2, 6})

	xor4, _ := xor.Base64("RWRnZXdoZXJl", "LnJvb290Lmlv")
	assert.DeepEqual(t, "axYIChgcSxsK", xor4)

	_, err := xor.Base64("ABCD", "ABC=")
	assert.Error(t, err, xor.NewNotSameLengthError().Error())

	notBase64 := "ABC;"
	_, err = xor.Base64(notBase64, "ABC=")
	assert.Error(t, err, "not a Base64-encoded string: "+notBase64)

	d1 := "test"
	keys := "keys"
	xor5, _ := xor.String(d1, keys)
	d2, _ := xor.String(xor5, keys)
	assert.Equal(t, d1, d2)
}

// TestComplement ...
func TestComplement(t *testing.T) {
	data := []byte("test")
	code := []byte("code")
	omega := xor.Complement(code)

	xor1, _ := xor.Bytes(data, omega)
	assert.DeepEqual(t, xor1, []byte{232, 245, 232, 238})
	deXorWrong, _ := xor.Bytes(xor1, code)
	assert.Assert(t, string(deXorWrong) != string(data))
	deXorRight, _ := xor.Bytes(xor1, omega)
	assert.Assert(t, string(deXorRight) == string(data))

	// The complement of two XORed items is not equal to the XOR of the their two complements
	plus := []byte("plus")
	codePlus, _ := xor.Bytes(code, plus)
	cpCompl := xor.Complement(codePlus)
	cpComplCompl, _ := xor.Bytes(xor.Complement(code), xor.Complement(plus))
	assert.Assert(t, string(cpCompl) != string(cpComplCompl))
}
