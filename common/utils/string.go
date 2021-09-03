package utils

import (
	"errors"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/charmap"
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

// Encoding formats
var (
	ISO_8859_3   = charmap.ISO8859_3.String()   // ISO 8859-3
	WINDOWS_1252 = charmap.Windows1252.String() // Windows 1252
)

// ToUTF8 transforms a encoded string in the passed format to a UTF-8 string
//
// NB: For now, it only supports ISO 8859-3 and Windows 1252 formats
func ToUTF8(encoded, format string) (decoded string, err error) {
	if format != ISO_8859_3 && format != WINDOWS_1252 {
		err = errors.New("unsupported format")
		return
	}
	foundReader, err := charset.NewReader(strings.NewReader(encoded), format)
	if err != nil {
		return
	}
	foundLine, err := ioutil.ReadAll(foundReader)
	if err != nil {
		return
	}
	decoded = string(foundLine)
	return
}
