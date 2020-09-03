package model_test

import (
	"testing"

	"github.com/cyrildever/go-utls/model"
	"gotest.tools/assert"
)

// TestSignatureBytes ...
func TestSignatureBytes(t *testing.T) {
	signature := model.Signature("181bfbf39202d365d812ffc608723dabb0d9438c018cd9ee6ecc209928f4a10114c47a7d49707961030dedaa564792c5ad06dc512f540173561e36edd1c7579000")
	bytes := []byte{24, 27, 251, 243, 146, 2, 211, 101, 216, 18, 255, 198, 8, 114, 61, 171, 176, 217, 67, 140, 1, 140, 217, 238, 110, 204, 32, 153, 40, 244, 161, 1, 20, 196, 122, 125, 73, 112, 121, 97, 3, 13, 237, 170, 86, 71, 146, 197, 173, 6, 220, 81, 47, 84, 1, 115, 86, 30, 54, 237, 209, 199, 87, 144, 0}
	assert.DeepEqual(t, bytes, signature.Bytes())
}

// TestSignatureString ...
func TestSignatureString(t *testing.T) {
	ref := "1234567890abcdef"
	sig := model.Signature(ref)
	assert.Equal(t, sig.String(), ref)

	wrong := model.Signature("wrong signature")
	assert.Assert(t, wrong.IsEmpty())
}

// TestToSignature ...
func TestToSignature(t *testing.T) {
	signature := model.Signature("181bfbf39202d365d812ffc608723dabb0d9438c018cd9ee6ecc209928f4a10114c47a7d49707961030dedaa564792c5ad06dc512f540173561e36edd1c7579000")

	bytes := []byte{24, 27, 251, 243, 146, 2, 211, 101, 216, 18, 255, 198, 8, 114, 61, 171, 176, 217, 67, 140, 1, 140, 217, 238, 110, 204, 32, 153, 40, 244, 161, 1, 20, 196, 122, 125, 73, 112, 121, 97, 3, 13, 237, 170, 86, 71, 146, 197, 173, 6, 220, 81, 47, 84, 1, 115, 86, 30, 54, 237, 209, 199, 87, 144, 0}
	sig := model.ToSignature(bytes)

	assert.DeepEqual(t, sig.Bytes(), signature.Bytes())
}
