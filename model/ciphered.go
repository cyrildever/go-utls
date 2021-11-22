package model

import (
	"encoding/base64"

	"github.com/cyrildever/go-utls/common/utils"
)

const (
	CIPHERED = "ciphered"
)

//--- TYPES

// Ciphered is the base64 string representation of a ciphered text.
type Ciphered string

//--- METHODS

// Bytes returns the underlying byte array if it's an actual base64-encoded string, or nil.
func (c Ciphered) Bytes() []byte {
	if c.String() == "" || !IsBase64String(c.String()) {
		return nil
	}
	return utils.Must(base64.StdEncoding.DecodeString(string(c)))
}

// String returns the ciphered string if it's an actual base64-encoded string, or an empty string.
func (c Ciphered) String() string {
	if !IsBase64String(string(c)) {
		return ""
	}
	return string(c)
}

// IsEmpty ...
func (c Ciphered) IsEmpty() bool {
	return c.Bytes() == nil
}

// NonEmpty ...
func (c Ciphered) NonEmpty() bool {
	return c.String() != ""
}

//--- FUNCTIONS

// ToCiphered ...
func ToCiphered(bytes []byte) Ciphered {
	if bytes == nil {
		return Ciphered("")
	}
	return Ciphered(base64.StdEncoding.EncodeToString(bytes))
}
