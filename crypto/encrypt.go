package crypto

import (
	"crypto/rand"

	"github.com/ethereum/go-ethereum/crypto/ecies"
)

// Encrypt ...
func Encrypt(msg, decompressedPublicKey []byte) ([]byte, error) {
	// As of the latest version, we assume the byte arrays are ECDSA-compatible keys
	pk, err := ECDSAPublicKeyFrom(decompressedPublicKey)
	if err != nil {
		return nil, err
	}
	pubkey := (*ecies.ImportECDSAPublic(&pk))

	return ecies.Encrypt(rand.Reader, &pubkey, msg, nil, nil)
}

// Decrypt ...
func Decrypt(ciphered, privateKey, decompressedPublicKey []byte) ([]byte, error) {
	// As of the latest version, we assume the byte arrays are ECDSA-compatible keys
	sk, err := ECDSAPrivateKeyFrom(privateKey, decompressedPublicKey)
	if err != nil {
		return nil, err
	}
	privkey := (*ecies.ImportECDSA(&sk))
	return privkey.Decrypt(ciphered, nil, nil)
}
