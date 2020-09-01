package utils

// Flatten ...
func Flatten(arrayOfByteArray [][]byte) (concat []byte) {
	for _, b := range arrayOfByteArray {
		concat = append(concat, b...)
	}
	return
}
