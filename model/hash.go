package model

import (
	"regexp"

	"github.com/cyrildever/go-utls/common/utils"
)

var hashRegex = regexp.MustCompile(`^[0-9a-fA-F]{32}([0-9a-fA-F]{32})?$`)

//--- TYPES

// Hash is the hexadecimal string representation of a hash.Hash
// It's either a 32-character or a 64-character long.
type Hash string

//--- METHODS

// Bytes returns the underlying byte array if it's an actual hash string, or nil.
func (h Hash) Bytes() []byte {
	if h.String() == "" || !CouldBeValidHash(h.String()) {
		return nil
	}
	return utils.Must(utils.FromHex(string(h)))
}

// String returns the hexadecimal string representation if it's an actual hash string, or an empty string.
func (h Hash) String() string {
	if !CouldBeValidHash(string(h)) {
		return ""
	}
	return string(h)
}

// IsEmpty ...
func (h Hash) IsEmpty() bool {
	return h.Bytes() == nil
}

// NonEmpty ...
func (h Hash) NonEmpty() bool {
	return h.String() != ""
}

//--- FUNCTIONS

// ToHash ...
func ToHash(bytes []byte) Hash {
	if bytes == nil {
		return Hash("")
	}
	return Hash(utils.ToHex(bytes))
}

// CouldBeValidHash returns `true` if the passed string could be a 32-bytes or 64-bytes hash in hexadecimal representation
func CouldBeValidHash(str string) bool {
	return hashRegex.MatchString(str)
}
