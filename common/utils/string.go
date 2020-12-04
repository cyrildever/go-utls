package utils

import (
	"strings"
)

// Capitalize add an upper case at the beginning of the sentence
func Capitalize(sentence string) string {
	if len(sentence) == 0 {
		return sentence
	}
	str := strings.ToUpper(sentence[:1])
	if len(sentence) > 1 {
		str += sentence[1:]
	}
	return str
}
