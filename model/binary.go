package model

import (
	"strconv"
	"strings"

	"github.com/cyrildever/go-utls/common/utils"
)

//--- TYPES

// Binary is the string representation of a binary literal, eg. Binary("1001") would equal to `0b1001` in other languages or starting with version 1.13 of Golang.
// NB: if the underlying string is not a valid string representation of a binary, it will not throw an error but rather return an empty or nil item for each method called.
type Binary string

//--- METHODS

// Bytes ...
func (b Binary) Bytes() []byte {
	if b.String() == "" {
		return nil
	}
	i, err := strconv.ParseUint(string(b), 2, 64)
	if err != nil {
		// This was not an appropriate binary literal in the first place
		return nil
	}
	barray := utils.UintToByteArray(i)
	return barray
}

// String ...
func (b Binary) String() string {
	if !IsBinaryString(string(b)) {
		return ""
	}
	return string(b)
}

// IsEmpty ...
func (b Binary) IsEmpty() bool {
	return b.Bytes() == nil
}

// NonEmpty ...
func (b Binary) NonEmpty() bool {
	return b.String() != ""
}

//--- FUNCTIONS

// ToBinary ...
func ToBinary(bytes []byte) Binary {
	if bytes == nil {
		return Binary("")
	}
	i := utils.ByteArrayToUint(bytes)
	if i == 0 {
		str := strings.Repeat("0", len(bytes))
		return Binary(str)
	}
	str := strconv.FormatUint(i, 2)
	if !IsBinaryString(str) {
		return Binary("")
	}
	return Binary(str)
}

// IsBinaryString ...
func IsBinaryString(str string) bool {
	if str == "" {
		return false
	}
	for _, c := range str {
		if c != '0' && c != '1' {
			return false
		}
	}
	return true
}
