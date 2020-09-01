package packer_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/cyrildever/go-utls/common/packer"
	"gotest.tools/assert"
)

// TestMarshal ...
func TestMarshal(t *testing.T) {
	ref := []byte{172, 123, 34, 116, 101, 115, 116, 34, 58, 49, 50, 51, 125}
	json := `{"test":123}`
	ret, err := packer.MessagePackMarshal(json)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(ret, ref) {
		fmt.Println(ret)
		t.Errorf("Bytes are not what they should be")
	}

	var jsonObject struct {
		Test int `json:"test"`
	}
	jsonObject.Test = 123
	found, err := packer.JSONMarshal(jsonObject)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, json, string(found))
}

// TestUnmarshal ...
func TestUnmarshal(t *testing.T) {
	ref := []byte{129, 164, 84, 101, 115, 116, 211, 0, 0, 0, 0, 0, 0, 0, 123}
	var jsonObject struct {
		Test int `json:"test"`
	}
	err := packer.MessagePackUnmarshal(ref, &jsonObject)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, jsonObject.Test, 123)

	json := `{"test":123}`
	err = packer.JSONUnmarshal([]byte(json), &jsonObject)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, jsonObject.Test, 123)
}
