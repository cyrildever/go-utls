package bip32

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/binary"
	"math"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/sammyne/base58"
)

// PrivateKey houses all the information of an extended private key.
type PrivateKey struct {
	PublicKey
	Data    []byte
	Version []byte
}

// AddressPubKeyHash implements ExtendedKey
func (priv *PrivateKey) AddressPubKeyHash() []byte {
	return btcutil.Hash160(priv.publicKeyData())
}

// Child implements ExtendedKey
func (priv *PrivateKey) Child(i uint32) (ExtendedKey, error) {
	// Prevent derivation of children beyond the max allowed depth.
	if priv.Level == math.MaxUint8 {
		return nil, ErrDeriveBeyondMaxDepth
	}

	// There are four scenarios that could happen here:
	// 1) Private extended key -> Hardened child private extended key
	// 2) Private extended key -> Non-hardened child private extended key
	// 3) Public extended key -> Non-hardened child public extended key
	// 4) Public extended key -> Hardened child public extended key (INVALID!)
	// where only #1, #2 is applicable in our context

	data := make([]byte, KeyDataLen+ChildIndexLen)
	if i < HardenedKeyStart { // normal
		copy(data, priv.publicKeyData())
	} else { // harden, where 0x00 prefix plus 32-byte data
		data[0] = 0x00
		ReverseCopy(data[1:KeyDataLen], priv.Data)
	}
	binary.BigEndian.PutUint32(data[KeyDataLen:], i)

	// Take the HMAC-SHA512 of the current key's chain code and the derived
	// data:
	//   I = HMAC-SHA512(Key = chainCode, Data = data)
	hmac512 := hmac.New(sha512.New, priv.ChainCode)
	hmac512.Write(data)
	I := hmac512.Sum(nil)

	IL, chainCode := I[:len(I)/2], I[len(I)/2:]

	// Both derived public or private keys rely on treating the left 32-byte
	// sequence calculated above (IL) as a 256-bit integer that must be
	// within the valid range for a secp256k1 private key.  There is a small
	// chance (< 1 in 2^127) this condition will not hold, and in that case,
	// a child extended key can't be created for this index and the caller
	// should simply increment to the next index.
	z, usable := ToUsableScalar(IL)
	if !usable {
		return nil, ErrInvalidChild
	}

	// Case #1 or #2.
	// Add the parent private key to the intermediate private key to
	// derive the final child key.
	//
	// childKey = parse256(IL) + parenKey
	k := new(big.Int).SetBytes(priv.Data)
	z.Add(z, k)
	z.Mod(z, secp256k1Curve.N)
	childData := z.Bytes()

	// The fingerprint of the parent for the derived child is the first 4
	// bytes of the RIPEMD160(SHA256(parentPubKey)).
	parentFP := priv.AddressPubKeyHash()[:FingerprintLen]

	return NewPrivateKey(priv.Version, priv.Level+1, parentFP, i,
		chainCode, childData), nil
}

// IsForNet implements ExtendedKey
func (priv *PrivateKey) IsForNet(keyID Magic) bool {
	return bytes.Equal(priv.Version, keyID[:])
}

// Neuter implements ExtendedKey
func (priv *PrivateKey) Neuter() (*PublicKey, error) {
	// Get the associated public extended key version bytes.
	version, err := chaincfg.HDPrivateKeyToPublicKeyID(priv.Version)
	if err != nil {
		return nil, err
	}

	/*
		// consider returns the internal public key instead to save allocation
		pub := priv.PublicKey // copy the common part

		// and update the different parts
		pub.Version = version

		data := priv.publicKeyData()
		pub.Data = make([]byte, len(data))
		copy(pub.Data, data)

		return &pub, nil
	*/
	priv.PublicKey.Version = version
	priv.PublicKey.Data = priv.publicKeyData()

	return &priv.PublicKey, nil
}

// Public implements ExtendedKey
func (priv *PrivateKey) Public() (*btcec.PublicKey, error) {
	return btcec.ParsePubKey(priv.publicKeyData(), secp256k1Curve)
}

// SetNet implements ExtendedKey
func (priv *PrivateKey) SetNet(keyID Magic) {
	priv.Version = keyID[:]
}

// String implements ExtendedKey
func (priv *PrivateKey) String() string {
	if 0 == len(priv.Data) {
		return "zeroed private key"
	}

	// The serialized format is:
	//   version (4) || depth (1) || parent fingerprint (4)) ||
	//   child num (4) || chain code (32) || key data (33)
	buf := make([]byte, 0, KeyLen-VersionLen)
	buf = appendMeta(buf, &priv.PublicKey)
	buf = append(buf, 0x00)
	buf = paddedAppend(KeyDataLen-1, buf, priv.Data)

	return base58.CheckEncodeX(buf, priv.Version...)
}

// ToECPrivate converts the extended key to a btcec private key and returns it.
// As you might imagine this is only possible if the extended key is a private
// extended key.
func (priv *PrivateKey) ToECPrivate() *btcec.PrivateKey {
	privKey, _ := btcec.PrivKeyFromBytes(secp256k1Curve, priv.Data)

	return privKey
}

// Zero implements ExtendedKey
func (priv *PrivateKey) Zero() {
	if nil == priv {
		return
	}

	priv.PublicKey.Zero()
	Zero(priv.Data)
	Zero(priv.Version)
}

// publicKeyData load the corresponding public key data in compressed form
// bound to this private key. It will check whether the data part of the public
// key is initialised, and initialise it if necessary.
func (priv *PrivateKey) publicKeyData() []byte {
	if 0 == len(priv.PublicKey.Data) {
		x, y := secp256k1Curve.ScalarBaseMult(priv.Data)
		pubKey := btcec.PublicKey{Curve: secp256k1Curve, X: x, Y: y}

		priv.PublicKey.Data = pubKey.SerializeCompressed()
	}

	return priv.PublicKey.Data
}

// NewPrivateKey returns a new instance of an extended private key with the
// given fields. No error checking is performed here as it's only intended to
// be a convenience method used to create a populated struct. This function
// should only by used by applications that need to create custom PrivateKey.
// All other applications should just use NewMasterKey, Child, or Neuter.
// **The public key part is left uninitialized yet**
func NewPrivateKey(version []byte, depth uint8, parentFP []byte, index uint32,
	chainCode, data []byte) *PrivateKey {
	pub := PublicKey{
		Level:      depth,
		ParentFP:   parentFP,
		ChildIndex: index,
		ChainCode:  chainCode,
	}

	priv := &PrivateKey{PublicKey: pub, Data: data, Version: version}

	// derive public key eagerly
	// this should be considered more seriously
	//priv.PublicKey.Data = derivePublicKey(priv.Data)

	return priv
}

// ParsePrivateKey a new extended private key instance out of a base58-encoded
// extended key.
// **The public key part is left uninitialized yet**
func ParsePrivateKey(data58 string) (*PrivateKey, error) {
	// decodePublicKey is applicable here too !!!
	pub, err := decodePublicKey(data58)
	if nil != err {
		return nil, err
	}

	priv := &PrivateKey{
		PublicKey: *pub,
		Data:      pub.Data[1:], // simply trims out the 0x00 prefix
	}
	priv.Version = priv.PublicKey.Version
	priv.PublicKey.Data, priv.PublicKey.Version = nil, nil

	// load the public key data eagerly
	/*
		x, y := secp256k1Curve.ScalarBaseMult(priv.Data)
		pubKey := btcec.PublicKey{Curve: secp256k1Curve, X: x, Y: y}
		priv.PublicKey.Data = pubKey.SerializeCompressed()
	*/
	//priv.PublicKey.Data = derivePublicKey(priv.Data)

	return priv, nil
}
