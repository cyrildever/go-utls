package crypto_test

import (
	"testing"

	"github.com/cyrildever/go-utls/crypto"
	"gotest.tools/assert"
)

// TestNext ...
func TestNext(t *testing.T) {
	ref := crypto.Path("m/1'/0/0")
	p := crypto.Path("m/0'/65535/2097151")
	next, err := p.Next()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, ref, next)
	nextOne, _ := next.Next()
	assert.Equal(t, nextOne.String(), "m/1'/0/1")
}

// TestParse ...
func TestParse(t *testing.T) {
	ref := crypto.Indices{
		Account: crypto.Account{
			Number:   0,
			Hardened: true,
		},
		Scope:    0,
		KeyIndex: 123,
	}

	p := crypto.Path("m/0'/0/123")
	indices, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, ref, indices)
}

// TestString ...
func TestString(t *testing.T) {
	defaut := "m/0'/0/0"
	assert.Equal(t, crypto.Path(defaut).String(), defaut)

	wrong := crypto.Path("not a path")
	assert.Assert(t, wrong.IsEmpty())
	assert.Equal(t, wrong.String(), "")
}
