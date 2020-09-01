package crypto_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/cyrildever/go-utls/common/utils"
	"github.com/cyrildever/go-utls/crypto"
	"gotest.tools/assert"
)

// TestValidKeyPair
func TestValidKeyPair(t *testing.T) {
	pk := utils.Must(utils.FromHex(crypto.PUBLIC_DECOMPRESSED_KEY_1_STRING))
	sk := utils.Must(utils.FromHex(crypto.SECRET_KEY_1_STRING))
	_, _, err := crypto.ECIESKeyPairFrom(pk, sk)
	assert.Assert(t, err == nil)
}

// TestInvalidKeyPair ...
func TestInvalidKeyPair(t *testing.T) {
	_, _, err := crypto.ECIESKeyPairFrom([]byte("invalid_key_1"), []byte(""))
	assert.Assert(t, err != nil)

	pk := utils.Must(utils.FromHex(crypto.PUBLIC_DECOMPRESSED_KEY_1_STRING))
	_, _, err = crypto.ECIESKeyPairFrom(pk, []byte("invalid_key_2"))
	assert.Assert(t, err != nil)
}

// TestKeypair ...
func TestKeypair(t *testing.T) {
	// From seed
	seed := utils.Must(utils.FromHex("e5dcc2bdf7310ea3abc82a3580621bd0a487c181113d889d0775d3f8eba21e84"))

	pk1, sk1, _ := crypto.GenerateKeyPair(seed, crypto.Path("m/0'/0/0"))
	pubkey1 := utils.ToHex(pk1)
	privkey1 := utils.ToHex(sk1)
	assert.Equal(t, pubkey1, "04e315a987bd79b9f49d3a1c8bd1ef5a401a242820d52a3f22505da81dfcd992cc5c6e2ae9bc0754856ca68652516551d46121daa37afc609036ab5754fe7a82a3")
	assert.Equal(t, privkey1, "b9fc3b425d6c1745b9c963c97e6e1d4c1db7a093a36e0cf7c0bf85dc1130b8a0")

	// From base58 private key
	pk2, _ := crypto.ParsePrivateKey("xprv9zMRtaKN7WvsSgEE7nFdUaCXhvJdCY31gSmYG93nSi5qhJt1UCbaWpiyGLFVZivz6QUCvQvXUew9y9D9PkVJCYzR9EBHYtjUJcJz7seneaB")
	pubkey2 := utils.ToHex(pk2)
	assert.Equal(t, pubkey1, pubkey2)

	// From base58 public key
	pk3, _ := crypto.ParsePublicKey("xpub6DLnJ5rFwtVAfAJhDondqi9GFx97bzks3fh94XTQ13cpa7DA1juq4d3T7dY71ay4guQbtwU8XQQji2YvWfYTyZiZM3XgYqCE5iAkSZeY8be")
	pubkey3 := utils.ToHex(pk3)
	assert.Equal(t, pubkey2, pubkey3)

	// From bytes
	pk4, _ := crypto.ParsePublicKeyFromCompressed("03e315a987bd79b9f49d3a1c8bd1ef5a401a242820d52a3f22505da81dfcd992cc")
	pubkey4 := utils.ToHex(pk4)
	assert.Equal(t, pubkey3, pubkey4)

	assert.Assert(t, crypto.IsCompatibleKeyPair(pk1, sk1))
	assert.Assert(t, crypto.IsCompatibleKeyPair(pk2, sk1))
	assert.Assert(t, crypto.IsCompatibleKeyPair(pk3, sk1))
	assert.Assert(t, crypto.IsCompatibleKeyPair(pk4, sk1))
}

// TestGenerateKeyPair ...
func TestGenerateKeyPair(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	seed, _ := crypto.Hash(utils.IntToByteArray(int(r.Int63())))
	fmt.Println("seed", utils.ToHex(seed))
	pk, sk, err := crypto.GenerateKeyPair(seed, "m/0'/0/0")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("pubKey", utils.ToHex(pk))
	fmt.Println("privKey", utils.ToHex(sk))
	assert.Assert(t, err == nil)
	// mk, _ := crypto.Hash(seed)
	// idkey, _ := crypto.Hash(mk)
	// fmt.Println("idkey", utils.ToHex(idkey))
	// assert.Assert(t, false)
}
