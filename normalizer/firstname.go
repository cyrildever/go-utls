package normalizer

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

var firstNameDictionary Dictionary

func init() {
	firstNameDictionary = make(map[string]string)

	file, err := os.Open("dictionary/firstname.dico")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.SplitN(line, "\t", 2)
		firstNameDictionary[items[0]] = items[1]
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}
}

// FirstName returns a normalized first name
var FirstName SimpleNormalizer = func(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("invalid empty string")
	}
	found := firstNameDictionary.TranslateWord(Uniformize(input))
	if found == "" {
		return "", errors.New("unable to build a normalized string")
	}
	return found, nil
}
