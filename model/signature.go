package model

import (
	"github.com/cyrildever/go-utls/common/utils"
)

//--- TYPES

// Signature is the hexadecimal string representation of a signature.
type Signature string

//--- METHODS

// Bytes ...
func (s Signature) Bytes() []byte {
	if s.String() == "" {
		return nil
	}
	return utils.Must(utils.FromHex(string(s)))
}

// String ...
func (s Signature) String() string {
	return string(s)
}

// IsEmpty ...
func (s Signature) IsEmpty() bool {
	return s.Bytes() == nil
}

// NonEmpty ...
func (s Signature) NonEmpty() bool {
	return s.String() != ""
}

//--- FUNCTIONS

// ToSignature ...
func ToSignature(bytes []byte) Signature {
	if bytes == nil {
		return Signature("")
	}
	return Signature(utils.ToHex(bytes))
}
