package crypto_test

import (
	"fmt"
	"testing"

	"github.com/cyrildever/go-utls/common/utils"
	"github.com/cyrildever/go-utls/crypto"
	"github.com/cyrildever/go-utls/model"
	"gotest.tools/assert"
)

var (
	pk = utils.Must(utils.FromHex(crypto.PUBLIC_DECOMPRESSED_KEY_1_STRING))
	sk = utils.Must(utils.FromHex(crypto.SECRET_KEY_1_STRING))
)

// TestEncrypt ...
func TestEncrypt(t *testing.T) {
	msg := []byte("Edgewhere")

	crypted, err := crypto.Encrypt(msg, pk)
	fmt.Println(model.ToBase64(crypted).String())
	if err != nil {
		t.Fatal(err)
	}
}

// TestDecrypt ...
func TestDecrypt(t *testing.T) {
	expected := []byte("Edgewhere")

	encrypted := model.Ciphered("BIRC2pLEQjDDZMoNjoHV8RXehmlMx8eBp+WWMKw0Uyr+FGl5NXv7drbu7INkQdal0L3D+sgV5Gh5Qj6hIltC4rSMAxZGA21ziciSl1wsVOPgPQbrIvEnT6trptT+8W58a3hJpJc9MMM5aoJrxDMVNiUJs7vkbq6SiRY=").Bytes()

	found, err := crypto.Decrypt(encrypted, sk, pk)
	if err != nil {
		t.Fatal(err)
	}
	assert.DeepEqual(t, expected, found)
}

// TestInvariance ...
func TestInvariance(t *testing.T) {
	expected := []byte("Edgewhere")

	encrypted, _ := crypto.Encrypt(expected, pk)
	found, _ := crypto.Decrypt(encrypted, sk, pk)

	assert.DeepEqual(t, found, expected)
}

// TestDifferentCiphers ...
func TestDifferentCiphers(t *testing.T) {
	value1 := model.Ciphered("BAAc2BBq/svTtOhs1Gpy45JfTGRCLm+QSJK06kNni3R17s5xKhE+TXR4s85Ag0dGZ3MH1/B6QBGhmtLGZ94+gjxm/bBgmrEqyW50fyhg2Bka7jOOGl2uyV9UVWQeL5wrxzgZPYt4uLAdkrGfQ2QRo5zlGxHvVrooxSg=")
	value2 := model.Ciphered("BDS+hSuDn6KdrVm2yRb+Eaj23gsBT+01ajc6HkD7j8tbHFQvy+4J+KbwtMUaQFdB46yYec9J/DKP1CxqFlO9IhsyMsrqJtDX8Z0S4u2NHquHL4HhBxNJ7/XBrkSzVYwOG5sWZUPo7B0mkBRn1ognLgtzmf7ed+GNFxw=")

	found1, _ := crypto.Decrypt(value1.Bytes(), sk, pk)
	found2, _ := crypto.Decrypt(value2.Bytes(), sk, pk)

	expected := []byte("Edgewhere")

	assert.DeepEqual(t, expected, found1)
	assert.DeepEqual(t, expected, found2)
}
