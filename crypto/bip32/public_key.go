package bip32

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/binary"
	"math"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
	"github.com/sammyne/base58"
)

// PublicKey is the structure layout for an extended public key.
type PublicKey struct {
	ChainCode  []byte
	ChildIndex uint32 // this is the Index-th child of its parent
	Data       []byte // the serialized data in the compressed form
	Level      uint8  // name so to avoid conflict with method Depth()
	ParentFP   []byte
	Version    []byte
}

// AddressPubKeyHash implements ExtendedKey
func (pub *PublicKey) AddressPubKeyHash() []byte {
	return btcutil.Hash160(pub.Data)
}

// Child derive a normal(non-hardened) child for current key
func (pub *PublicKey) Child(i uint32) (ExtendedKey, error) {
	// Prevent derivation of children beyond the max allowed depth.
	if pub.Level == math.MaxUint8 {
		return nil, ErrDeriveBeyondMaxDepth
	}

	// There are four scenarios that could happen here:
	// 1) Private extended key -> Hardened child private extended key
	// 2) Private extended key -> Non-hardened child private extended key
	// 3) Public extended key -> Non-hardened child public extended key
	// 4) Public extended key -> Hardened child public extended key (INVALID!)
	// where only #3 is applicable in our context

	if i >= HardenedKeyStart {
		return nil, ErrDeriveHardFromPublic
	}

	// concatenate data and index
	data := make([]byte, KeyDataLen+4)
	copy(data, pub.Data)
	binary.BigEndian.PutUint32(data[KeyDataLen:], i)

	// Take the HMAC-SHA512 of the current key's chain code and the derived
	// data:
	//   I = HMAC-SHA512(Key = chainCode, Data = data)
	hmac512 := hmac.New(sha512.New, pub.ChainCode)
	hmac512.Write(data)
	I := hmac512.Sum(nil)

	IL, chainCode := I[:len(I)/2], I[len(I)/2:]

	// Both derived public or private keys rely on treating the left 32-byte
	// sequence calculated above (IL) as a 256-bit integer that must be
	// within the valid range for a secp256k1 private key.  There is a small
	// chance (< 1 in 2^127) this condition will not hold, and in that case,
	// a child extended key can't be created for this index and the caller
	// should simply increment to the next index.
	if _, ok := ToUsableScalar(IL); !ok {
		return nil, ErrInvalidChild
	}

	// Calculate the corresponding intermediate public key for
	// intermediate private key.
	x, y := btcec.S256().ScalarBaseMult(IL)
	if x.Sign() == 0 || y.Sign() == 0 {
		return nil, ErrInvalidChild
	}

	// Convert the serialized compressed parent public key into X
	// and Y coordinates so it can be added to the intermediate
	// public key.
	Kp, err := btcec.ParsePubKey(pub.Data, secp256k1Curve)
	if err != nil {
		return nil, err
	}

	// Add the intermediate public key to the parent public key to
	// derive the final child key.
	//
	// childKey = serP(point(parse256(I_L)) + K_par)
	x, y = secp256k1Curve.Add(x, y, Kp.X, Kp.Y)
	child := btcec.PublicKey{Curve: secp256k1Curve, X: x, Y: y}
	childData := child.SerializeCompressed()

	FP := btcutil.Hash160(pub.Data)[:FingerprintLen]
	return NewPublicKey(pub.Version, pub.Level+1, FP, i, chainCode,
		childData), nil
}

// Depth implements ExtendedKey
func (pub *PublicKey) Depth() uint8 {
	return pub.Level
}

// Hardened implements ExtendedKey
func (pub *PublicKey) Hardened() bool {
	return pub.ChildIndex >= HardenedKeyStart
}

// Index implements ExtendedKey
func (pub *PublicKey) Index() uint32 {
	return pub.ChildIndex
}

// IsForNet implements ExtendedKey
func (pub *PublicKey) IsForNet(keyID Magic) bool {
	return bytes.Equal(pub.Version, keyID[:])
}

// Neuter implements ExtendedKey
func (pub *PublicKey) Neuter() (*PublicKey, error) {
	return pub, nil
}

// ParentFingerprint implements ExtendedKey
func (pub *PublicKey) ParentFingerprint() uint32 {
	//panic("not implemented")
	return binary.BigEndian.Uint32(pub.ParentFP)
}

// Public implements ExtendedKey
func (pub *PublicKey) Public() (*btcec.PublicKey, error) {
	return btcec.ParsePubKey(pub.Data, secp256k1Curve)
}

// SetNet implements ExtendedKey
func (pub *PublicKey) SetNet(keyID Magic) {
	pub.Version = keyID[:]
}

// String implements ExtendedKey
func (pub *PublicKey) String() string {
	if 0 == len(pub.Data) {
		return "zeroed public key"
	}

	var childIndex [ChildIndexLen]byte
	binary.BigEndian.PutUint32(childIndex[:], pub.ChildIndex)

	// The serialized format is:
	//   version (4) || depth (1) || parent fingerprint (4)) ||
	//   child num (4) || chain code (32) || key data (33)
	str := make([]byte, 0, KeyLen-VersionLen)
	str = append(str, pub.Level)
	str = append(str, pub.ParentFP...)
	str = append(str, childIndex[:]...)
	str = append(str, pub.ChainCode...)
	str = append(str, pub.Data...)

	return base58.CheckEncodeX(str, pub.Version...)
}

// Zero implements ExtendedKey
func (pub *PublicKey) Zero() {
	if nil == pub {
		return
	}

	Zero(pub.ChainCode)
	pub.ChildIndex = 0
	Zero(pub.Data)
	pub.Level = 0
	Zero(pub.ParentFP)
	Zero(pub.Version)
}

// NewPublicKey returns a new instance of an extended public key with the
// given fields. No error checking is performed here as it's only intended to
// be a convenience method used to create a populated struct. This function
// should only by used by applications that need to create custom PublicKey.
// All other applications should just use Child or Neuter.
func NewPublicKey(version []byte, depth uint8, parentFP []byte, index uint32,
	chainCode, data []byte) *PublicKey {
	return &PublicKey{
		Version:    version,
		Level:      depth,
		ParentFP:   parentFP,
		ChildIndex: index,
		ChainCode:  chainCode,
		Data:       data,
	}
}

// ParsePublicKey a new extended public key instance out of a base58-encoded
// extended key.
func ParsePublicKey(data58 string) (*PublicKey, error) {
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

	// on-curve checking
	if _, err := btcec.ParsePubKey(pub.Data, secp256k1Curve); nil != err {
		return nil, err
	}

	return pub, nil
}
