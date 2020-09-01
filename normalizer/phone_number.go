package normalizer

import (
	"errors"
	"regexp"
	"strings"
)

var reTel = regexp.MustCompile(`^(((00)?(33)|0?(262)|0?(590)|0?(594)|0?(596))([^1-9]*)|0)?([^0-9]*)([\d]{3})([\d]{3})([\d]{3})$`)

// PhoneNumber returns a normalized landline phone number, or an empty string and an error if failed.
// WARNING: The current implementation is specific to French mobile phones
// TODO Become international
var PhoneNumber SimpleNormalizer = func(input string) (string, error) {
	matches := reTel.FindStringSubmatch(strings.ReplaceAll(Uniformize(input), " ", ""))
	if len(matches) < 1 {
		return "", errors.New("invalid phone number")
	}
	var parts = make([]string, 5)
	var p1, p2, p3, p4, p5 bool
	for i, v := range matches {
		switch i {
		case 4:
			{
				international := "+"
				if len(v) > 0 {
					international += v
				} else {
					international += "33"
				}
				parts[0] = international
				p1 = true
			}
		case 9:
			{
				prefix := "("
				if len(v) > 0 {
					prefix += v
				} else {
					prefix += "0"
				}
				prefix += ")"
				parts[1] = prefix
				p2 = true
			}
		case 11:
			if v != "" {
				parts[2] = v
				p3 = true
			}
		case 12:
			if v != "" {
				parts[3] = v
				p4 = true
			}
		case 13:
			if v != "" {
				parts[4] = v
				p5 = true
			}
		default:
			// NO-OP
		}
	}
	if !p1 || !p2 || !p3 || !p4 || !p5 {
		return "", errors.New("unable to build normalized phone number")
	}
	return strings.Join(parts, " "), nil
}
