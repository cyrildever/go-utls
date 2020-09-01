package utils

import (
	"encoding/binary"
)

// UintToByteArray converts the passed integer into a byte array.
func UintToByteArray(data uint64) (barray []byte) {
	barray = make([]byte, 8)
	binary.LittleEndian.PutUint64(barray, data)
	return
}

// ByteArrayToUint converts the passed byte array into a uint64.
func ByteArrayToUint(barray []byte) uint64 {
	return binary.LittleEndian.Uint64(barray)
}

// Uint8ToByteArray converts the passed uint8 into a byte array.
func Uint8ToByteArray(data uint8) (barray []byte) {
	barray = make([]byte, 1)
	barray[0] = byte(data)
	return
}

// ByteArrayToUint8 converts the passed byte array into a uint8.
func ByteArrayToUint8(barray []byte) uint8 {
	return uint8(barray[0])
}

// IntToByteArray converts the passed integer to a byte array.
func IntToByteArray(data int) []byte {
	return UintToByteArray(uint64(data))
}

// ByteArrayToInt converts the passed byte array to a int.
func ByteArrayToInt(bytes []byte) int {
	return int(binary.LittleEndian.Uint64(bytes))
}
