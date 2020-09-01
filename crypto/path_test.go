package crypto_test

import (
	"testing"

	"github.com/cyrildever/go-utls/crypto"
	"gotest.tools/assert"
)

// TestParse ...
func TestParse(t *testing.T) {
	ref := crypto.Indices{
		Account: crypto.Account{
			Number:   0,
			Hardened: true,
		},
		Scope:    0,
		KeyIndex: 0,
	}

	p := crypto.Path("m/0'/0/0")
	indices, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, ref, indices)
}
