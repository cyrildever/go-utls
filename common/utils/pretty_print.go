package utils

import (
	"bytes"
	"encoding/json"

	"github.com/cyrildever/go-utls/common/packer"
)

// PrettyPrintJSON ...
func PrettyPrintJSON(input interface{}) ([]byte, error) {
	b, err := packer.JSONMarshal(input)
	if err != nil {
		return nil, err
	}
	return PrettyPrintByte(b)
}

// PrettyPrintByte ...
func PrettyPrintByte(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}
