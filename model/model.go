package model

// Model defines the basic interface for each core model type.
// All types should also define their own To<TypeName> function to get an object of the type from a byte array,
// eg. `func ToHash(bytes []byte) Hash` for the `Hash` type.
type Model interface {
	Bytes() []byte
	String() string
	NonEmpty() bool
}
