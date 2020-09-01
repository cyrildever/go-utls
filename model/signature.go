package model

import (
	"github.com/cyrildever/go-utls/common/utils"
)

// Signature is the hexadecimal string representation of a signature.
type Signature string

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

// NonEmpty ...
func (s Signature) NonEmpty() bool {
	return s.String() != ""
}

// ToSignature ...
func ToSignature(bytes []byte) Signature {
	if bytes == nil {
		return Signature("")
	}
	return Signature(utils.ToHex(bytes))
}
