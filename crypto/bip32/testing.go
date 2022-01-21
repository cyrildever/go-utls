package bip32

import (
	"encoding/hex"
	"io"
	"strconv"
	"strings"
)

// This file specify some utility structure to ease testing.

//go:generate go run golden.go

// GoldenBase base directory storing the golden files
const GoldenBase = "testdata"

// GoldenName specifies the file containing the reference test vectors
const GoldenName = "bip32.golden"

// GoldenAddOnName specifies the file containing the more test vectors
// from the btcutil/hdkeychain testing
const GoldenAddOnName = "bip32_addon.golden"

// ChainGoldie is the structure for key chains
type ChainGoldie struct {
	Path               Path // child stemming from master node
	ExtendedPublicKey  string
	ExtendedPrivateKey string
}

// ChildIndex is enhanced index capable of telling whether its hardened.
type ChildIndex struct {
	Index    uint32
	Hardened bool
}

// Goldie is the structure of a single test vector.
type Goldie struct {
	Seed   string
	Chains []ChainGoldie
}

// Path alias string with a path decoding api
type Path string

// ChildIndices return the child indices tree along this path
func (path Path) ChildIndices() ([]*ChildIndex, error) {
	indices := strings.Split(string(path), "/")

	childs := make([]*ChildIndex, len(indices)-1)
	for i, v := range indices[1:] { // skip the root
		hardened := strings.HasSuffix(v, "H")

		index, err := strconv.Atoi(strings.TrimSuffix(v, "H"))
		if nil != err {
			return nil, err
		}

		childs[i] = &ChildIndex{Index: uint32(index), Hardened: hardened}
	}

	return childs, nil
}

// NewEntropyReader constructs hex decoder for testing.
func NewEntropyReader(hexStr string) io.Reader {
	return hex.NewDecoder(strings.NewReader(hexStr))
}
