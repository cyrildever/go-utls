package crypto

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
	"github.com/cyrildever/go-utls/common/utils"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/sammyne/base58"
)

const (
	//--- ERROR MESSAGES
	invalidKeysErrorTxt = "invalid EC keys"
)

var secp256k1Curve = btcec.S256()

// BIP32PublicKey is the structure layout for an extended public key.
type BIP32PublicKey struct {
	ChainCode  []byte
	ChildIndex uint32 // this is the Index-th child of its parent
	Data       []byte // the serialized data in the compressed form
	Level      uint8  // name so to avoid conflict with method Depth()
	ParentFP   []byte
	Version    []byte
}

// BIP32PrivateKey houses all the information of an extended private key.
type BIP32PrivateKey struct {
	BIP32PublicKey
	Data    []byte
	Version []byte
}

// GenerateKeyPair is the function that should be used whenever a keypair is needed in the Rooot project.
//
// It takes a seed (ie. the master key in Rooot) and a BIP-32 path, and returns a keypair in byte array, decompressed public and private,
// and an error if any.
// As of the latest version in the system, this keypair operates on Bitcoin's secp256k1 elliptic curve.
// It shall be used for either signing with ECDSA or encrypting/decrypting with ECIES.
// One should use the utils.ToHex() utility method to save any of these byte arrays as their hexadecimal string representation if needed.
func GenerateKeyPair(seed []byte, path Path) (pubkey, privkey []byte, err error) {
	wallet, err := hdwallet.NewFromSeed(seed)
	if err != nil {
		return
	}
	p := hdwallet.MustParseDerivationPath(path.String())
	acc, err := wallet.Derive(p, false)
	if err != nil {
		return
	}
	sk, err := wallet.PrivateKeyBytes(acc)
	if err != nil {
		return
	}
	pk, err := wallet.PublicKeyBytes(acc)
	if err != nil {
		return
	}
	pubkey = pk
	privkey = sk
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
	// decodePublicKey is applicable here too !!!
	pub, err := decodeKey(base58PrivateKey)
	if nil != err {
		return nil, err
	}

	priv := &BIP32PrivateKey{
		BIP32PublicKey: *pub,
		Data:           pub.Data[1:], // simply trims out the 0x00 prefix
	}
	priv.Version = priv.BIP32PublicKey.Version
	priv.BIP32PublicKey.Data, priv.BIP32PublicKey.Version = nil, nil

	return derivePublicKeyFrom(priv)
}

// ParsePublicKey ...
func ParsePublicKey(base58PublicKey string) (pubkey []byte, err error) {
	pk, err := decodeKey(base58PublicKey, true)
	if err != nil {
		return
	}
	pub, err := btcec.ParsePubKey(pk.Data, secp256k1Curve)
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

func decodeKey(data58 string, isPublic ...bool) (pub *BIP32PublicKey, err error) {
	decoded, version, err := base58.CheckDecodeX(data58, 4)
	if err != nil {
		return
	}
	if len(decoded)+4 != 78 {
		return nil, fmt.Errorf("invalid key length for %s", data58)
	}

	pk := new(BIP32PublicKey)

	// Decompose the decoded payload into fields
	//
	// The serialized format is:
	//   version (4) || depth (1) || parent fingerprint (4)) ||
	//   child num (4) || chain code (32) || key data (33)
	// where the version was already separated from decoded
	pk.Version = version

	a, b := 0, 1
	pk.Level = decoded[a:b][0]

	a, b = b, b+4
	pk.ParentFP = decoded[a:b]

	a, b = b, b+4
	pk.ChildIndex = binary.BigEndian.Uint32(decoded[a:b])

	a, b = b, b+32
	pk.ChainCode = decoded[a:b]

	a, b = b, b+33
	pk.Data = decoded[a:b]

	// On-curve checking for actual public keys
	if len(isPublic) == 1 && isPublic[0] {
		if _, err := btcec.ParsePubKey(pk.Data, secp256k1Curve); err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	pub = pk
	return
}

func derivePublicKeyFrom(key *BIP32PrivateKey) (pubkey []byte, err error) {
	if 0 == len(key.BIP32PublicKey.Data) {
		x, y := secp256k1Curve.ScalarBaseMult(key.Data)
		pubKey := btcec.PublicKey{Curve: secp256k1Curve, X: x, Y: y}

		key.BIP32PublicKey.Data = pubKey.SerializeCompressed()
	}
	pk, err := btcec.ParsePubKey(key.BIP32PublicKey.Data, secp256k1Curve)
	if err != nil {
		return
	}
	pubkey = elliptic.Marshal(secp256k1.S256(), pk.X, pk.Y)
	return
}
