package crypto

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/cyrildever/go-utls/common/utils"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/sammyne/bip32"
)

const (
	//--- ERROR MESSAGES
	invalidKeysErrorTxt = "invalid EC keys"
)

// GenerateKeyPair is the function that should be used whenever a keypair is needed in the Rooot project.
//
// It takes a seed (ie. the master key in Rooot) and a BIP-32 path, and returns a keypair in byte array, decompressed public and private,
// and an error if any.
// As of the latest version in the system, this keypair operates on Bitcoin's secp256k1 elliptic curve.
// It shall be used for either signing with ECDSA or encrypting/decrypting with ECIES.
// One should use the utils.ToHex() utility method to save any of these byte arrays as their hexadecimal string representation if needed.
func GenerateKeyPair(seed []byte, path Path) (pubkey, privkey []byte, err error) {
	mk, err := bip32.NewMasterKey(seed, *bip32.MainNetPrivateKey)
	if err != nil {
		return
	}
	indices, err := path.Parse()
	if err != nil {
		return
	}
	var hk bip32.ExtendedKey
	if indices.Account.Hardened {
		hk, err = mk.Child(bip32.HardenIndex(indices.Account.Number))
		if err != nil {
			return
		}
	} else {
		hk, err = mk.Child(indices.Account.Number)
		if err != nil {
			return
		}
	}
	ek, err := hk.Child(indices.Scope)
	if err != nil {
		return
	}
	xk, err := ek.Child(indices.KeyIndex)
	if err != nil {
		return
	}
	pubkey, err = derivePublicKeyFrom(xk.(*bip32.PrivateKey))
	if err != nil {
		return
	}
	newPrivate, err := bip32.ParsePrivateKey(xk.String())
	privkey = newPrivate.ToECPrivate().Serialize()
	return
}

// GenerateRandomKeyPair generates a random keypair
func GenerateRandomKeyPair() (pubkey, privkey []byte, err error) {
	key, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		return
	}
	pubkey = elliptic.Marshal(secp256k1.S256(), key.X, key.Y)

	privkey = make([]byte, 32)
	blob := key.D.Bytes()
	copy(privkey[32-len(blob):], blob)

	return
}

// ParsePrivateKey ...
func ParsePrivateKey(base58PrivateKey string) (pubkey []byte, err error) {
	sk, err := bip32.ParsePrivateKey(base58PrivateKey)
	if err != nil {
		return
	}
	return derivePublicKeyFrom(sk)
}

// ParsePublicKey ...
func ParsePublicKey(base58PublicKey string) (pubkey []byte, err error) {
	pk, err := bip32.ParsePublicKey(base58PublicKey)
	if err != nil {
		return
	}
	pub, err := pk.Public()
	if err != nil {
		return
	}
	pubkey = elliptic.Marshal(secp256k1.S256(), pub.X, pub.Y)
	return
}

// ParsePublicKeyFromCompressed ...
func ParsePublicKeyFromCompressed(str string) (pubkey []byte, err error) {
	bytes, err := utils.FromHex(str)
	if err != nil {
		return
	}
	x, y := secp256k1.DecompressPubkey(bytes)
	pubkey = elliptic.Marshal(secp256k1.S256(), x, y)
	return
}

// ECIESKeyPairFrom transforms an ECDSA keypair in byte array to its ECIES counterparts.
//
// It takes a keypair in byte array and returns the ECIES PublicKey and PrivateKey objects, and an error if something went wrong.
// Keys generated through the GenerateKeyPair() function should be compatible.
func ECIESKeyPairFrom(ecdsaPublicKBytes, ecdsaPrivateKBytes []byte) (pubkey ecies.PublicKey, privkey ecies.PrivateKey, err error) {
	// Public Key
	pk, err := ECDSAPublicKeyFrom(ecdsaPublicKBytes)
	if err != nil {
		return
	}
	pubkey = (*ecies.ImportECDSAPublic(&pk))

	// Private Key
	sk, err := ECDSAPrivateKeyFrom(ecdsaPrivateKBytes, ecdsaPublicKBytes)
	if err != nil {
		return
	}
	privkey = (*ecies.ImportECDSA(&sk))
	if !privkey.IsOnCurve(privkey.X, privkey.Y) {
		err = fmt.Errorf("%s: invalid private key", invalidKeysErrorTxt)
		return
	}

	if !IsCompatibleKeyPair(ecdsaPublicKBytes, ecdsaPrivateKBytes) {
		err = fmt.Errorf("%s: incompatible keys", invalidKeysErrorTxt)
		return
	}

	return pubkey, privkey, err
}

// ECDSAPublicKeyFrom ...
func ECDSAPublicKeyFrom(publicKeyBytes []byte) (pk ecdsa.PublicKey, err error) {
	x, y := elliptic.Unmarshal(secp256k1.S256(), publicKeyBytes)
	if x == nil || y == nil {
		err = fmt.Errorf("%s: invalid ECDSA public key", invalidKeysErrorTxt)
		return
	}
	pk = ecdsa.PublicKey{
		Curve: secp256k1.S256(),
		X:     x,
		Y:     y,
	}
	return
}

// ECDSAPrivateKeyFrom ...
func ECDSAPrivateKeyFrom(privateKeyBytes, publicKeyBytes []byte) (sk ecdsa.PrivateKey, err error) {
	pk, err := ECDSAPublicKeyFrom(publicKeyBytes)
	if err != nil {
		return
	}
	d := new(big.Int)
	d.SetBytes(privateKeyBytes)
	sk = ecdsa.PrivateKey{
		PublicKey: pk,
		D:         d,
	}
	return
}

// IsCompatibleKeyPair tests whether a decompressed public key and a private key form an actual valid keypair.
func IsCompatibleKeyPair(pk, sk []byte) bool {
	msgTest := []byte{255}
	ct, err := Encrypt(msgTest, pk)
	if err != nil {
		return false
	}
	dt, err := Decrypt(ct, sk, pk)
	if err != nil {
		return false
	}
	if !bytes.Equal(msgTest, dt) {
		return false
	}
	return true
}

// --- utility functions

func derivePublicKeyFrom(xk *bip32.PrivateKey) (pubkey []byte, err error) {
	pk, err := xk.Public()
	if err != nil {
		return
	}
	pubkey = elliptic.Marshal(secp256k1.S256(), pk.X, pk.Y)
	return
}
