package model

import (
	"encoding/base64"

	"github.com/cyrildever/go-utls/common/utils"
)

//--- TYPES

// Ciphered is the base64 string representation of a ciphered text.
type Ciphered string

//--- METHODS

// Bytes ...
func (c Ciphered) Bytes() []byte {
	if c.String() == "" {
		return nil
	}
	return utils.Must(base64.StdEncoding.DecodeString(string(c)))
}

// String ...
func (c Ciphered) String() string {
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
