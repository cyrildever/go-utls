package utils

// Flatten takes an array of byte array and makes it a byte array
func Flatten(arrayOfByteArray [][]byte) (concat []byte) {
	for _, b := range arrayOfByteArray {
		concat = append(concat, b...)
	}
	return
}
