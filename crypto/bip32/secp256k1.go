package bip32

import (
	"math/big"

	"github.com/btcsuite/btcd/btcec"
)

var secp256k1Curve = btcec.S256()

// ToUsableScalar tries to covert k to a scalar 0<s<secp256k1Curve.N as a big
// integer, and returns false in case of failure.
func ToUsableScalar(k []byte) (*big.Int, bool) {
	x := new(big.Int).SetBytes(k)

	return x, x.Cmp(secp256k1Curve.N) < 0 && 0 != x.Sign()
}
