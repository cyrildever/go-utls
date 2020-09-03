package model

import (
	"errors"
	"strings"
)

//--- TYPES

// Model defines the basic interface for each core model type.
//
// All types should also define their own `To<TypeName>` function to get an object of the type from a byte array,
// eg. `func ToHash(bytes []byte) Hash` for the `Hash` type.
//
// Invalid data for the `Model` type should not throw errors but rather act as if they were empty.
type Model interface {
	// Bytes returns the underlying byte array of the data
	Bytes() []byte

	// String returns the usual string representation of the data
	String() string

	// IsEmpty checks whether there is no data, ie. it's an empty byte array
	IsEmpty() bool

	// NonEmpty tells if the underlying data is not void
	NonEmpty() bool
}

//--- FUNCTIONS

// ToModel is a factory utility.
// It takes the byte array of the data and the name of the type to return.
func ToModel(bytes []byte, typeName string) (m Model, err error) {
	// TODO Enrich for each new type
	switch strings.ToLower(typeName) {
	case "base64":
		m = ToBase64(bytes)
	case "binary":
		m = ToBinary(bytes)
	case "bits8":
		m = ToBits8(bytes)
	case "ciphered":
		m = ToCiphered(bytes)
	case "hash":
		m = ToHash(bytes)
	case "key":
		m = ToKey(bytes)
	case "signature":
		m = ToSignature(bytes)
	case "uuid":
		m = ToUUID(bytes)
	default:
		err = errors.New("Invalid type")
	}
	return
}
