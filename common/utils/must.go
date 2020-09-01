package utils

// Must only returns the byte array from the argument tuple (byte array, error).
//
// Use with caution, ie. only when you are absolutely certain that it shouldn't lead to a panic,
// or that if it leads to a panic it means that there's no way the program should work anyway.
// For example, `bytes := utils.Must(hex.DecodeString(h))` where `h` is an hexadecimal string hash.
func Must(bytes []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return bytes
}
