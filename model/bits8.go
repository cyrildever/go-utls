package model

import (
	"errors"

	"github.com/cyrildever/go-utls/common/utils"
)

const (
	BITS_8 = "bits8"
)

//--- TYPES

// Bits8 is an 8-bit representation of a bit set, ie. a single byte/octet
type Bits8 uint8

//--- METHODS

// Bytes ...
func (b Bits8) Bytes() []byte {
	return utils.Uint8ToByteArray(uint8(b))
}

// String ...
func (b Bits8) String() string {
	var str string
	for i := 0; i < 8; i++ {
		flag := 1
		flag = flag << uint(i)
		if b.Has(Bits8(flag)) {
			str = "1" + str
		} else {
			str = "0" + str
		}
	}
	return str
}

// IsEmpty ...
func (b Bits8) IsEmpty() bool {
	return b.Bytes() == nil
}

// NonEmpty ...
func (b Bits8) NonEmpty() bool {
	return b.String() != ""
}

// Set ...
func (b *Bits8) Set(flag Bits8) {
	bits := *b | flag
	*b = bits
}

// Clear ...
func (b *Bits8) Clear(flag Bits8) {
	bits := *b &^ flag
	*b = bits
}

// ClearAll ...
func (b *Bits8) ClearAll() {
	*b = *new(Bits8)
}

// Toggle ...
func (b *Bits8) Toggle(flag Bits8) {
	bits := *b ^ flag
	*b = bits
}

// Has ...
func (b *Bits8) Has(flag Bits8) bool {
	return *b&flag != 0
}

//--- FUNCTIONS

// ToBits8 ...
func ToBits8(bytes []byte) Bits8 {
	return Bits8(utils.ByteArrayToUint8(bytes))
}

// StringToBits8 ...
func StringToBits8(str string) (b Bits8, err error) {
	if len(str) > 8 {
		err = errors.New("input string too long to be a valid 8-bit representation")
		return
	}
	// Add padding if necessary
	if len(str) < 8 {
		for i := len(str); i < 8; i++ {
			str = "0" + str
		}
	}
	for i, c := range str {
		var flag Bits8
		if c == '1' {
			flag = 1
			flag = flag << uint(8-i-1)
			b.Set(flag)
		} else if c != '0' {
			err = errors.New("not an 8-bit representation")
			return
		}
	}
	return
}

// IsBits8String ...
func IsBits8String(str string) bool {
	if str == "" || len(str) != 8 {
		return false
	}
	for _, c := range str {
		if c != '0' && c != '1' {
			return false
		}
	}
	return true
}
