package normalizer

import (
	"strings"
)

//--- TYPES

// Dictionary ...
type Dictionary map[string]string

//--- METHODS

// TranslateText returns the passed string ith all its words passed through the TranslateWord function.
// The input should have been uniformized beforehand.
func (d Dictionary) TranslateText(input string) string {
	var translated []string
	words := strings.Split(input, " ")
	for _, word := range words {
		tmp := d.TranslateWord(word)
		if tmp != "" {
			translated = append(translated, tmp)
		}
	}
	return strings.Join(translated, " ")
}

// TranslateWord returns the translated word if found in the dictionary, the original string otherwise.
// The input should have been uniformized beforehand.
func (d Dictionary) TranslateWord(input string) string {
	if translated, ok := d[input]; ok {
		return translated
	}
	return input
}

//--- SHARED
var addressDictionary Dictionary
