package model

import (
	"regexp"
)

var hexRegex = regexp.MustCompile(`^([0-9a-fA-F]{2})+$`)

//--- FUNCTIONS

// IsHexString checks whether the passed string is an hexadecimal number of even length.
func IsHexString(str string) bool {
	return hexRegex.MatchString(str)
}
