package model

import (
	"reflect"
	"regexp"

	"github.com/cyrildever/go-utls/common/utils"
)

// Hash is the hexadecimal string representation of a hash.
type Hash string

// Bytes ...
func (h Hash) Bytes() []byte {
	if h.String() == "" {
		return nil
	}
	return utils.Must(utils.FromHex(string(h)))
}

// String ...
func (h Hash) String() string {
	return string(h)
}

// NonEmpty ...
func (h Hash) NonEmpty() bool {
	return h.String() != ""
}

// ToHash ...
func ToHash(bytes []byte) Hash {
	if bytes == nil {
		return Hash("")
	}
	return Hash(utils.ToHex(bytes))
}

var hashRegex = regexp.MustCompile(`^[0-9a-fA-F]{32}([0-9a-fA-F]{32})?$`)

// CouldBeValidHash returns `true` if the passed string could be a 32-bytes or 64-bytes hash in hexadecimal representation
func CouldBeValidHash(str string) bool {
	return hashRegex.MatchString(str)
}

// Hashes is an array of Hash.
type Hashes []Hash

func (h Hashes) Len() int           { return len(h) }
func (h Hashes) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Hashes) Less(i, j int) bool { return h[i] < h[j] }

// Equals ...
func (h *Hashes) Equals(to *Hashes) bool {
	length := h.Len()
	if length != to.Len() {
		return false
	}
	for i := 0; i < length; i++ {
		if !reflect.DeepEqual((*h)[i], (*to)[i]) {
			return false
		}
	}
	return true
}

// Contains ...
func (h Hashes) Contains(item Hash) bool {
	for _, hash := range h {
		if hash == item {
			return true
		}
	}
	return false
}
