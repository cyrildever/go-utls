package crypto_test

import (
	"testing"

	"github.com/cyrildever/go-utls/common/utils"
	"github.com/cyrildever/go-utls/crypto"
	"github.com/cyrildever/go-utls/model"
	"gotest.tools/assert"
)

// TestSign ...
func TestSign(t *testing.T) {
	sk, _ := utils.FromHex(crypto.SECRET_KEY_1_STRING)
	sig1, _ := utils.FromHex("963cb6e1036b5355333b5f4f491ef075664db00fbc98037c4d1a69e5cbd950822c4d6a4e1ef203ab0a4b04d0fcb4f70ed8de74bf565c01067a55c18bb7401a87")

	msg1 := []byte("Edgewhere")
	signature, err := crypto.Sign(msg1, sk)
	if err != nil {
		t.Fatal(err)
	}
	assert.DeepEqual(t, sig1, signature)
}

// TestCheck ...
func TestCheck(t *testing.T) {
	pk, _ := utils.FromHex(crypto.PUBLIC_DECOMPRESSED_KEY_1_STRING)
	sig1, _ := utils.FromHex("963cb6e1036b5355333b5f4f491ef075664db00fbc98037c4d1a69e5cbd950822c4d6a4e1ef203ab0a4b04d0fcb4f70ed8de74bf565c01067a55c18bb7401a87")
	msg1 := []byte("Edgewhere")
	assert.Assert(t, crypto.Check(sig1, msg1, pk))

	concat := []byte("opt-in")
	concat = append(concat, []byte("email")...)
	start := utils.IntToByteArray(309936)
	concat = append(concat, start...)
	end := utils.IntToByteArray(3153600)
	concat = append(concat, end...)
	maxUse := utils.IntToByteArray(1)
	concat = append(concat, maxUse...)
	purposes := "0x7f"
	concat = append(concat, []byte(purposes)...)
	scopes := "0xf10xAAZZDghYBV5SWVgOCg4§CQMGUVNeB1wEDwMNDwI"
	concat = append(concat, []byte(scopes)...)

	pkB, _ := utils.FromHex("04417ee29bdd3b7bdcb1ecd0cb20e3e704afacca2e6be75b88220f9639d207960ce9473f1750f4cafb56e24ada51aa4f7dbe301880837fab61ceb56f8715bb68a3")
	sigB, _ := utils.FromHex("fdc46d2ea2cba88ff2df26f3c657a82768f974de6b45a91fa71b77ed06cdb1064db997976e8942c624edcb976ec3c634654ecafb57fb7fe426dfd90786d1b51a")
	pkE, _ := utils.FromHex("040c96f971c0edf58fe4afbf8735581be05554a8a725eae2b7ad2b1c6fcb7b39ef4e7252ed5b17940a9201c089bf75cb11f97e5c53333a424e4ebcca36065e0bc0") // It's the good one
	sigE, _ := utils.FromHex("0eafe75bcd0389c0cd4163f2b5653bb282b60a827114c7a84d53bfa406eae41b0971bf28f66af64a45fc32669763e43d2ca735da9606a94760ca34ed68e8a982")
	assert.Assert(t, crypto.Check(sigB, concat, pkB) || crypto.Check(sigE, concat, pkE))
}

// TestSignAndCheck ...
func TestSignAndCheck(t *testing.T) {
	msg := []byte("Edgewhere")

	pk1, sk1, _ := crypto.GenerateRandomKeyPair()
	signature, err := crypto.Sign(msg, sk1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Assert(t, crypto.Check(signature, msg, pk1))

	pk2, sk2, _ := crypto.GenerateKeyPair(utils.Must(utils.FromHex(crypto.SEED_TEST_STRING)), crypto.PATH_1_STRING)
	signature, err = crypto.Sign(msg, sk2)
	if err != nil {
		t.Fatal(err)
	}
	assert.Assert(t, crypto.Check(signature, msg, pk2))
}

// TestEciesGethSignatures ...
func TestEciesGethSignatures(t *testing.T) {
	transactionId := model.Hash("682c2b9eb3fd843cb9d9294756769e81974e75921c06f43cc7f66476895183de")
	signedHashedTransactionID := model.Base64("MEQCIGXWjzHuf6Sx3mnvcqRUo+vcjOcjjpATuQpUcCNl4yHcAiBa+reQJQOANQRolxtyQGYCYpaHzIKqqUWUlDB61T/tBQ==")
	publicKey := model.Key("04788c1cb5ade3a0843b54df45aa777b0182b6570121bb577c82b63a815a4265a68aff6d06b621099083f2c66365e027a3fd94b25c40af43dda7d86adea9958910")
	assert.Assert(t, crypto.Check(signedHashedTransactionID.Bytes(), transactionId.Bytes(), publicKey.Bytes()))

	alreadyValidSig := []byte{101, 214, 143, 49, 238, 127, 164, 177, 222, 105, 239, 114, 164, 84, 163, 235, 220, 140, 231, 35, 142, 144, 19, 185, 10, 84, 112, 35, 101, 227, 33, 220, 90, 250, 183, 144, 37, 3, 128, 53, 4, 104, 151, 27, 114, 64, 102, 2, 98, 150, 135, 204, 130, 170, 169, 69, 148, 148, 48, 122, 213, 63, 237, 5}
	signature, err := crypto.ImportSignature(signedHashedTransactionID.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	assert.DeepEqual(t, alreadyValidSig, signature)
	assert.Assert(t, crypto.Check(alreadyValidSig, transactionId.Bytes(), publicKey.Bytes()))
}
