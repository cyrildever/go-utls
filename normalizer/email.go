package normalizer

import (
	"errors"
	"regexp"
	"strings"
)

var reEmail = regexp.MustCompile("^(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f]))@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+(?:[a-z]{2,}[a-z0-9-]*|([a-z][a-z0-9-]+))|(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.?){4})$")

// Email returns a normalized e-mail
var Email SimpleNormalizer = func(input string) (string, error) {
	// see https://emailregex.com/
	processed := strings.ToLower(strings.TrimSpace(input))
	if len(processed) > 255 || !reEmail.MatchString(processed) {
		return "", errors.New("invalid email")
	}
	return processed, nil
}
