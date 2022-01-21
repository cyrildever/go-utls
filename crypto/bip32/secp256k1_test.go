package bip32_test

import (
	"math/big"
	"testing"

	"github.com/btcsuite/btcd/btcec"
	"github.com/cyrildever/go-utls/crypto/bip32"
)

func TestToUsableScalar(t *testing.T) {
	type expect struct {
		s  *big.Int
		ok bool
	}

	testCases := []struct {
		k      []byte
		expect expect
	}{
		{ // a valid k
			[]byte{0x12, 0x34, 0x56, 0x78},
			expect{
				new(big.Int).SetInt64(0x12345678),
				true,
			},
		},
		{ // zero scalar
			[]byte{0x00},
			expect{nil, false},
		},
		{ // two large scalar
			btcec.S256().N.Bytes(),
			expect{nil, false},
		},
	}

	for i, c := range testCases {
		got, ok := bip32.ToUsableScalar(c.k)

		if ok != c.expect.ok {
			t.Fatalf("#%d invalid status: got %v, expect %v", i, ok, c.expect.ok)
		}

		if ok && 0 != c.expect.s.Cmp(got) {
			t.Fatalf("#%d invalid scalar: got %v, expect %v", i, got, c.expect.s)
		}
	}
}
