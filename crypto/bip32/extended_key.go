package bip32

import "github.com/btcsuite/btcd/btcec"

// References:
//   [BIP32]: BIP0032 - Hierarchical Deterministic Wallets
//   https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki

// ExtendedKey specifies the basic api for a extended private or public key.
type ExtendedKey interface {
	// AddressPubKeyHash returns the public key hash (20 bytes) derived from the
	// key, which would be itself for public keys and the extended public key
	// bound to it for private keys
	AddressPubKeyHash() []byte
	// Child returns a derived child extended key at the given index.  When this
	// extended key is a private extended key, a private extended key will be
	// derived.  Otherwise, the derived extended key will be also be a public
	// extended key.
	//
	// When the index is greater to or equal than the HardenedKeyStart constant,
	// the derived extended key will be a hardened extended key.  It is only
	// possible to derive a hardended extended key from a private extended key.
	// Consequently, this function will return ErrDeriveHardFromPublic if a
	// hardened child extended key is requested from a public extended key.
	//
	// A hardened extended key is useful since, as previously mentioned, it
	// requires a parent private extended key to derive. In other words, normal
	// child extended public keys can be derived from a parent public extended
	// key (no knowledge of the parent private key) whereas hardened extended
	// keys may not be.
	//
	// NOTE: There is an extremely small chance (< 1 in 2^127) the specific child
	// index does not derive to a usable child.  The ErrInvalidChild error should
	// be returned if this should occur, and the caller is expected to ignore the
	// invalid child and simply increment to the next index.
	Child(i uint32) (ExtendedKey, error)
	// Depth returns the current derivation level with respect to the root.
	//
	// The root key has depth zero, and the field has a maximum of 255 due to
	// how depth is serialized.
	Depth() uint8
	// Hardened tells if the key of interest is hardened, whose index is greater
	// than HardenedKeyStart.
	Hardened() bool
	// Index returns the index of the key seen from its parent.
	Index() uint32
	// IsForNet checks if this key is in use for a given net specified by keyID
	IsForNet(keyID Magic) bool

	// Neuter returns a new extended public key from a extended private key. The
	// same extended key will be returned unaltered if it is already an extended
	// public key.
	//
	// As the name implies, an extended public key does not have access to the
	// private key, so it is not capable of signing transactions or deriving
	// child extended private keys.  However, it is capable of deriving further
	// child extended public keys.
	Neuter() (*PublicKey, error)
	// ParentFingerprint returns a fingerprint of the parent extended key from
	// which this one was derived.
	//
	// It's defined the be the first 32 bits of the key identifier as specified by
	// https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#key-identifiers
	ParentFingerprint() uint32
	// Public converts the extended key to a btcec public key and returns it.
	Public() (*btcec.PublicKey, error)
	// SetNet associates the extended key, and any child keys yet to be derived
	// from it, with the passed network.
	SetNet(keyID Magic)
	// String returns the extended key as a human-readable base58-encoded string.
	String() string

	// Zero manually clears all fields and bytes in the
	// extended key. This can be used to explicitly clear key material from memory
	// for enhanced security against memory scraping. This function only clears
	// this particular key and not any children that have already been derived.
	Zero()
}
