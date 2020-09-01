package crypto

import (
	"errors"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

// Sign signs a content using secp256k1 elliptic curve.
//
// It takes the content to sign and the secret key to use as arguments and returns the signature and an error if any.
// As of the latest version of the system, it first hashes the content message with the system's Hash() function,
// then signs it using Bitcoin's secp256k1 elliptic curve that returns a recoverable ECDSA signature in 65-byte [R||S||V] format,
// and finally gets rid of the recovery 65th byte (the V).
func Sign(msg []byte, secretKey []byte) ([]byte, error) {
	hash, err := Hash(msg)
	if err != nil {
		return nil, err
	}
	rsvSignature, err := secp256k1.Sign(hash, secretKey)
	if err != nil {
		return nil, err
	}
	signature := rsvSignature[:64]
	return signature, nil
}

// Check verifies the validity of the signature of a given content with a public key.
//
// It takes the signature, the content that was signed and the decompressed public key to use as arguments and returns true if it checks out.
// It's the verifying counterpart of the Sig() function.
// Therefore, as of the latest version, it takes an ECDSA signature in 64-byte [R||S] format and its associated decompressed public key in byte array
// and hashes the content message with the system's Hash() function before using it.
func Check(signature []byte, msg []byte, decompressedPublicKey []byte) bool {
	if len(signature) == 0 {
		return false
	}
	sig, err := ImportSignature(signature)
	if err != nil {
		return false
	}
	hash, err := Hash(msg)
	if err != nil {
		return false
	}
	return secp256k1.VerifySignature(decompressedPublicKey, hash, sig)
}

// ImportSignature parses a DER ECDSA signature as in `ecies-geth` JavaScript module to produce the appropriate 64-byte [R||S] signature
// @see https://github.com/cryptocoinjs/secp256k1-node/blob/90a04a2e1127f4c1bfd7015aa5a7b22d08edb811/lib/elliptic.js
func ImportSignature(input []byte) (signature []byte, err error) {
	if len(input) == 64 {
		// Probably already a valid [R||S] signature
		return input, nil
	}
	invalidRS := errors.New("not a valid [R||S] signature")
	lenR := int(input[3])
	if lenR == 0 || 5+lenR > len(input) || input[4+lenR] != 2 {
		err = invalidRS
		return
	}
	lenS := int(input[5+lenR])
	if lenS == 0 || 6+lenR+lenS != len(input) {
		err = invalidRS
		return
	}

	sigR := input[4 : 4+lenR]
	if len(sigR) == 33 && sigR[0] == 0 {
		sigR = sigR[1:]
	}
	if len(sigR) > 32 {
		err = invalidRS
		return
	}

	sigS := input[6+lenR:]
	if len(sigS) == 33 && sigS[0] == 0 {
		sigS = sigS[1:]
	}
	if len(sigS) > 32 {
		err = invalidRS
		return
	}
	var finalSigS []byte
	if len(sigR)+len(sigS) < 64 {
		finalSigS = append([]byte{0}, sigS...)
	} else {
		finalSigS = sigS
	}

	sig := append(sigR, finalSigS...)
	if len(sig) != 64 {
		err = errors.New("invalid input signature")
		return
	}
	signature = sig
	return
}
