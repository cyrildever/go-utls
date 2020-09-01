package model_test

import (
	"testing"

	"github.com/cyrildever/go-utls/model"
	"gotest.tools/assert"
)

// TestUUID ...
func TestUUID(t *testing.T) {
	ref := "12345678-90ab-cdef-1234-67890abcdef1"
	rightUUID := model.UUID(ref)
	assert.Equal(t, rightUUID.String(), ref)

	wrongUUID := model.ToUUID([]byte("this-is-a-wrong-uuid"))
	assert.Equal(t, wrongUUID.String(), "")

	generated, err := model.GenerateUUID()
	if err != nil {
		t.Fatal(err)
	}
	assert.Assert(t, generated.NonEmpty())
	// fmt.Println(generated)
	// assert.Assert(t, false)
}
