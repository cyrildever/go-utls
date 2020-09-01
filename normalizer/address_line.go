package normalizer

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func init() {
	if len(addressDictionary) == 0 {
		addressDictionary = make(map[string]string)

		f, err := os.Open("dictionary/address.dico")
		if err != nil {
			return
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
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

// AddressLine returns a normalized address line
// (for line 2 to 6 in French postal convention but not only)
// TODO Become international in the address.dico file
var AddressLine SimpleNormalizer = func(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("invalid empty string")
	}
	found := addressDictionary.TranslateText(Uniformize(input))
	if found == "" {
		return "", errors.New("unable to build a normalized string")
	}
	return found, nil
}
