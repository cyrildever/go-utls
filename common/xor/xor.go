package xor

import (
	"fmt"

	"github.com/cyrildever/go-utls/model"
)

// Base64 function XOR two Base64-encoded string and returns the Base64-encoded result string
func Base64(item1, item2 string) (string, error) {
	if !model.IsBase64String(item1) {
		return "", fmt.Errorf("not a Base64-encoded string: %s", item1)
	}
	if !model.IsBase64String(item2) {
		return "", fmt.Errorf("not a Base64-encoded string: %s", item2)
	}
	b64_1 := model.Base64(item1)
	b64_2 := model.Base64(item2)
	bytes, err := Bytes(b64_1.Bytes(), b64_2.Bytes())
	if err != nil {
		return "", err
	}
	return model.ToBase64(bytes).String(), nil
}

// Bytes function XOR two byte arrays
func Bytes(item1, item2 []byte) ([]byte, error) {
	if len(item1) != len(item2) {
		return nil, NewNotSameLengthError()
	}
	n := len(item1)
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = item1[i] ^ item2[i]
	}
	return buf, nil
}

// String function XOR two strings in the sense that each charCode/rune are xored
func String(item1, item2 string) (string, error) {
	buf, err := Bytes([]byte(item1), []byte(item2))
	return string(buf), err
}

// Complement computes the XOR complement of the passed byte array, ie. Complement(x) = 1 ^ x
func Complement(item []byte) []byte {
	var complement = make([]byte, len(item))
	for i := 0; i < len(item); i++ {
		complement[i] = ^item[i]
	}
	return complement
}

// NotSameLengthError ...
type NotSameLengthError struct {
	message string
}

func (e NotSameLengthError) Error() string {
	return e.message
}

// NewNotSameLengthError ...
func NewNotSameLengthError() *NotSameLengthError {
	return &NotSameLengthError{
		message: "items are not of the same length",
	}
}
