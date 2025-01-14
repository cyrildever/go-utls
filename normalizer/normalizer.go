package normalizer

import (
	"errors"
	"io"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// The normalizer module is the Golang version of TypeScript's `es-normalizer` library and of the patented Empreinte Sociométrique's normalizing process.

// SimpleNormalizer ...
type SimpleNormalizer func(string) (string, error)

// VariadicNormalizer ...
type VariadicNormalizer func(string, ...string) (string, error)

// Normalize is the main function for the adaptation of the normalizing process
// developed for the Empreinte Sociometrique&trade; by Edgewhere:
// it takes the data to normalize and the normalizing function to use as well as some optional parameters,
// and returns the normalized string and an error if any.
func Normalize(data string, with interface{}, params ...string) (string, error) {
	if ptrfunc, ok := with.(VariadicNormalizer); ok {
		return ptrfunc(data, params...)
	}
	if ptrfunc, ok := with.(SimpleNormalizer); ok {
		return ptrfunc(data)
	}
	return "", errors.New("unknown normalizer type")
}

var re1 = regexp.MustCompile(`[^a-zA-Z0-9-]+`)
var re2 = regexp.MustCompile(`[-]+`)

// Uniformize applies the basic normalizing process: trim, capitalize, ...
func Uniformize(input string) string {
	transformer := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	normalized, err := io.ReadAll(transform.NewReader(strings.NewReader(strings.ToLower(strings.TrimSpace(input))), transformer))
	if err != nil {
		return ""
	}
	return strings.ToUpper(
		strings.TrimSpace(
			strings.ReplaceAll(
				re2.ReplaceAllString(
					re1.ReplaceAllString(
						strings.NewReplacer("ß", "ss", "ø", "o").Replace(string(normalized)),
						"-"),
					"-"),
				"-", " "),
		),
	)
}

// Any is the normalizing function to use as normalizer argument for any kind of data when no specific normalizer exists
var Any SimpleNormalizer = func(input string) (string, error) {
	uniformized := Uniformize(input)
	if uniformized == "" && input != "" {
		return "", errors.New("unable to normalize input string")
	}
	return uniformized, nil
}
