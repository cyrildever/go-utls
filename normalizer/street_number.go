package normalizer

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

var reSN = regexp.MustCompile(`^(\d*)\s*(.*)$`)

func init() {
	if len(addressDictionary) == 0 {
		addressDictionary = make(map[string]string)

		file, err := os.Open("dictionary/address.dico")
		if err != nil {
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			items := strings.SplitN(line, "\t", 2)
			addressDictionary[items[0]] = items[1]
		}

		if err = scanner.Err(); err != nil {
			panic(err)
		}
	}
}

// StreetNumber returns a sanitized street number
var StreetNumber SimpleNormalizer = func(input string) (string, error) {
	matches := reSN.FindAllStringSubmatch(strings.TrimSpace(input), 2)
	if len(matches) == 0 || len(matches[0]) != 3 {
		return "", errors.New("probably not a street number")
	}
	num := matches[0][1]
	compUni := ""
	if matches[0][2] != "" {
		comp := Uniformize(matches[0][2])
		if comp != "" {
			compUni = addressDictionary.TranslateText(comp)
		}
	}
	return num + compUni, nil
}
