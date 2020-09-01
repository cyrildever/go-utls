package normalizer

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

var reCity = regexp.MustCompile(`CDX(\s+[\d]+)*`)

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

// City returns a normalized city name
var City SimpleNormalizer = func(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("invalid empty string")
	}
	found := addressDictionary.TranslateText(Uniformize(input))
	if found == "" {
		return "", errors.New("unable to build a normalized string")
	}
	spaces := regexp.MustCompile(`\s+`)
	found = spaces.ReplaceAllString(reCity.ReplaceAllString(found, ""), " ")
	return strings.TrimSpace(found), nil
}
