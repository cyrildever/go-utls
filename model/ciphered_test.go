package model_test

import (
	"testing"

	"github.com/cyrildever/go-utls/model"
	"gotest.tools/assert"
)

// TestCipheredBytes ...
func TestCipheredBytes(t *testing.T) {
	ciphered := model.Ciphered("widrWaqcakeUnalpIfGe")
	bytes := []byte{194, 39, 107, 89, 170, 156, 106, 71, 148, 157, 169, 105, 33, 241, 158}
	assert.DeepEqual(t, bytes, ciphered.Bytes())
}

// TestToCiphered
func TestToCiphered(t *testing.T) {
	ref := model.Ciphered("widrWaqcakeUnalpIfGe")
	bytes := []byte{194, 39, 107, 89, 170, 156, 106, 71, 148, 157, 169, 105, 33, 241, 158}
	ciphered := model.ToCiphered(bytes)
	assert.DeepEqual(t, ref.Bytes(), ciphered.Bytes())
}
