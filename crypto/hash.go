package crypto

import (
	"crypto/sha256"

	"github.com/cyrildever/go-utls/common/utils"
)

// Hash is the hashing function that should be used whenever hashing in the rooot&trade; project.
//
// It takes a byte array as argument and returns a byte array and an error if any.
// As of the latest version of the system, it double-hashes any input string with the SHA256 standard algorithm.
// One should use the utils.ToHex() utility method to save the byte array as its hexadecimal string representation.
func Hash(input []byte) ([]byte, error) {
	h, err := hash(input)
	if err == nil {
		return hash(h)
	}
	return nil, err
}

// IsHashedValue ...
func IsHashedValue(str string) bool {
	if len(str) != 64 {
		return false // As long as we are using SHA256 algorithm
	}
	if _, err := utils.FromHex(str); err != nil {
		return false
	}
	return true
}

func hash(input []byte) ([]byte, error) {
	hasher := sha256.New()
	_, err := hasher.Write(input)
	return hasher.Sum(nil), err
}
