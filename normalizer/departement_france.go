package normalizer

import (
	"errors"
	"regexp"
	"strconv"
)

var reDF = regexp.MustCompile(`^(0[1-9]|[1-8][0-9]|9[0-5]|2[AB]|97[1-8]|98[46-9])([0-9]*)$`)

// DepartementFrance returns the two-letter code of a French "département"
var DepartementFrance SimpleNormalizer = func(input string) (string, error) {
	uniformized := Uniformize(input)
	matches := reDF.FindAllStringSubmatch(uniformized, 2)
	if len(matches) == 0 || len(matches[0]) != 3 {
		return "", errors.New("not a valid departement")
	}
	dpt := matches[0][1]
	if dpt == "20" {
		if cp, e := CodePostalFrance(input); e == nil {
			cpInt, _ := strconv.ParseInt(cp, 10, 64)
			if cpInt > 19999 && cpInt < 20200 {
				dpt = "2A"
			} else if cpInt > 20199 && cpInt < 20621 {
				dpt = "2B"
			}
		}
	}
	return dpt, nil
}
