package bip32

// ReverseCopy works in the opposite direction as the built-in copy()
func ReverseCopy(dst, src []byte) {
	for i, j := len(dst)-1, len(src)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		dst[i] = src[j]
	}
}

// paddedAppend appends the src byte slice to dst, returning the new slice.
// If the length of the source is smaller than the passed size, leading zero
// bytes are appended to the dst slice before appending src.
func paddedAppend(size uint, dst, src []byte) []byte {
	for i := 0; i < int(size)-len(src); i++ {
		dst = append(dst, 0)
	}
	return append(dst, src...)
}
