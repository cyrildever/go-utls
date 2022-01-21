package bip32

import (
	"encoding/binary"

	"github.com/btcsuite/btcd/btcec"
	"github.com/sammyne/base58"
)

// appendMeta serialize the meta part of the given public key
// into bytes sequence and then append them to the buf, the address
// of which will be return.
func appendMeta(buf []byte, pub *PublicKey) []byte {
	var childIndex [ChildIndexLen]byte
	binary.BigEndian.PutUint32(childIndex[:], pub.ChildIndex)

	// The serialized format of meta is:
	// depth (1) || parent fingerprint (4)) || child num (4) || chain code (32)
	// note the missing version and data fields

	//str := make([]byte, 0, KeyLen-VersionLen)
	//str = append(str, pub.Version...)
	buf = append(buf, pub.Level)
	buf = append(buf, pub.ParentFP...)
	buf = append(buf, childIndex[:]...)
	buf = append(buf, pub.ChainCode...)

	return buf
}

// decodePublicKey decodes a public key out of the given base58-check encoded
// key string.
// Note: the decoded key goes through format check only, no on-curve check
func decodePublicKey(data58 string) (*PublicKey, error) {
	decoded, version, err := base58.CheckDecodeX(data58, VersionLen)
	if nil != err {
		return nil, err
	}

	if KeyLen != len(decoded)+VersionLen {
		return nil, ErrInvalidKeyLen
	}

	pub := new(PublicKey)
	// The serialized format is:
	//   version (4) || depth (1) || parent fingerprint (4)) ||
	//   child num (4) || chain code (32) || key data (33)
	// where the version has separated from decoded

	// decompose the decoded payload into fields
	pub.Version = version

	a, b := 0, DepthLen
	pub.Level = decoded[a:b][0]

	a, b = b, b+FingerprintLen
	pub.ParentFP = decoded[a:b]

	a, b = b, b+ChildIndexLen
	pub.ChildIndex = binary.BigEndian.Uint32(decoded[a:b])

	a, b = b, b+ChainCodeLen
	pub.ChainCode = decoded[a:b]

	a, b = b, b+KeyDataLen
	pub.Data = decoded[a:b]

	return pub, nil
}

// derivePublicKey calculates the public key corresponding to the input
// private key, and output the public key data in compressed form.
func derivePublicKey(priv []byte) []byte {
	// load the public key data eagerly
	x, y := secp256k1Curve.ScalarBaseMult(priv)
	pubKey := btcec.PublicKey{Curve: secp256k1Curve, X: x, Y: y}

	return pubKey.SerializeCompressed()
}
