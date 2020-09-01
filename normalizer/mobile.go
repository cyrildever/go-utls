package normalizer

import (
	"errors"
	"regexp"
	"strings"
)

var reMob = regexp.MustCompile(`^(((00)?(33))|0)?([0]?)([\d])([\d]{2})([\d]{3})([\d]{3})$`)

// Mobile returns a normalized mobile, or an empty string and an error if failed.
// WARNING: The current implementation is specific to French mobile phones
// TODO Become international
var Mobile SimpleNormalizer = func(input string) (string, error) {
	matches := reMob.FindStringSubmatch(strings.ReplaceAll(Uniformize(input), " ", ""))
	if len(matches) < 1 {
		return "", errors.New("invalid mobile string")
	}
	var parts = make([]string, 5)
	var p1, p2, p3, p4, p5 bool
	mob := ""
	for i, v := range matches {
		if i == 4 {
			international := "+"
			if len(v) > 0 {
				international += v
			} else {
				international += "33"
			}
			parts[0] = international
			p1 = true
		}
		if i == 5 {
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
		if i == 6 {
			if v == "6" || v == "7" {
				mob = v
			} else {
				return "", errors.New("not a mobile phone")
			}
		}
		if i == 7 && mob != "" && v != "" {
			parts[2] = mob + v
			p3 = true
		}
		if i == 8 && v != "" {
			parts[3] = v
			p4 = true
		}
		if i == 9 && v != "" {
			parts[4] = v
			p5 = true
		}
	}
	if !p1 || !p2 || !p3 || !p4 || !p5 {
		return "", errors.New("unable to build normalized mobile")
	}
	return strings.Join(parts, " "), nil
}
