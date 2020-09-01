package normalizer

import (
	"errors"
	"regexp"
	"strings"
)

var reCPF = regexp.MustCompile(`((0[1-9][0-9]|[1-8][0-9]{2})|(9[0-5][0-9])|(2[AB][0-9])|(97[1-6]))[0-9]{2}`)
var reCorse = regexp.MustCompile(`^2[AB][0-9]{3}$`)

// CodePostalFrance returns a normalized French zip code
var CodePostalFrance SimpleNormalizer = func(input string) (string, error) {
	uniformized := Uniformize(input)
	if !reCPF.MatchString(uniformized) {
		return "", errors.New("invalid code postal")
	}
	if reCorse.MatchString(uniformized) {
		uniformized = strings.NewReplacer("A", "0", "B", "0").Replace(uniformized)
	}
	return uniformized, nil
}
