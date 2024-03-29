package model

import (
	"github.com/cyrildever/go-utls/common/utils"
)

const (
	KEY = "key"
)

//--- TYPES

// Key is the hexadecimal string representation of a public or private key.
type Key string

//--- METHODS

// Bytes ...
func (k Key) Bytes() []byte {
	if k.String() == "" || !IsHexString(string(k)) {
		return nil
	}
	return utils.Must(utils.FromHex(string(k)))
}

// String ...
func (k Key) String() string {
	if !IsHexString(string(k)) {
		return ""
	}
	return string(k)
}

// IsEmpty ...
func (k Key) IsEmpty() bool {
	return k.Bytes() == nil
}

// NonEmpty ...
func (k Key) NonEmpty() bool {
	return k.String() != ""
}

//--- FUNCTIONS

// ToKey ...
func ToKey(bytes []byte) Key {
	if bytes == nil {
		return Key("")
	}
	return Key(utils.ToHex(bytes))
}
