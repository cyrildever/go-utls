package utils

import (
	"errors"
	"io/ioutil"
	"strings"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/charmap"
)

// Encoding formats
var (
	ISO_8859_1   = charmap.ISO8859_1.String() // ISO 8859-1
	ISO_8859_2   = charmap.ISO8859_2.String() // ISO 8859-2
	ISO_8859_3   = charmap.ISO8859_3.String() // ISO 8859-3
	ISO_8859_4   = charmap.ISO8859_4.String() // ISO 8859-4
	MACINTOSH    = charmap.Macintosh.String() // Macintosh
	UTF_8        = "utf-8"
	WINDOWS_1252 = charmap.Windows1252.String() // Windows 1252
)

// ToUTF8 transforms a encoded string in the passed encoding format to a UTF-8 string
//
// NB: For now, it only supports ISO 8859-1, ISO 8859-2, ISO 8859-3, ISO 8859-4, Macintosh and Windows 1252 formats
func ToUTF8(encoded, format string) (decoded string, err error) {
	if strings.ToLower(format) == UTF_8 {
		return encoded, nil
	}
	if format != ISO_8859_1 &&
		format != ISO_8859_2 &&
		format != ISO_8859_3 &&
		format != ISO_8859_4 &&
		format != MACINTOSH &&
		format != WINDOWS_1252 {
		err = errors.New("unsupported encoding")
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
