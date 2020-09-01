package model

import (
	"github.com/cyrildever/go-utls/common/utils"
)

// Key is the hexadecimal string representation of a public or private key.
type Key string

// Bytes ...
func (k Key) Bytes() []byte {
	if k.String() == "" {
		return nil
	}
	return utils.Must(utils.FromHex(string(k)))
}

// String ...
func (k Key) String() string {
	return string(k)
}

// NonEmpty ...
func (k Key) NonEmpty() bool {
	return k.String() != ""
}

// ToKey ...
func ToKey(bytes []byte) Key {
	if bytes == nil {
		return Key("")
	}
	return Key(utils.ToHex(bytes))
}

// Keys is an array of Key.
type Keys []Key

// Contains ...
func (k Keys) Contains(item Key) bool {
	for _, key := range k {
		if key == item {
			return true
		}
	}
	return false
}
