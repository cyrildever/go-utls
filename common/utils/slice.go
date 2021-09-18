package utils

// Chunk splits a slice of strings into the number of passed chunks
func Chunk(slice []string, chunkSize int) [][]string {
	if len(slice) == 0 {
		return nil
	}
	chunks := make([][]string, (len(slice)+chunkSize-1)/chunkSize)
	prev := 0
	i := 0
	till := len(slice) - chunkSize
	for prev < till {
		next := prev + chunkSize
		chunks[i] = slice[prev:next]
		prev = next
		i++
	}
	chunks[i] = slice[prev:]
	return chunks
}

// ContainString returns true if the passed slice of strings contains the passed string value
func ContainString(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
