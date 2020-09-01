package normalizer

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

var reCiv = regexp.MustCompile(`\d+`)

var titleDictionary Dictionary

const (
	// CODE_F for female
	CODE_F = "2"

	// CODE_M for male
	CODE_M = "1"

	// CODE_U for unknown or undefined
	CODE_U = "0"
)

func init() {
	titleDictionary = make(map[string]string)

	file, err := os.Open("dictionary/title.dico")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.SplitN(line, "\t", 2)
		titleDictionary[items[0]] = items[1]
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}
}

// Title returns a code string: `0` for undefined, `1` for male, `2` for female
var Title SimpleNormalizer = func(input string) (string, error) {
	if len(input) == 0 {
		return "", errors.New("invalid empty string")
	}
	uniformized := Uniformize(input)
	if uniformized == "M" || uniformized == "H" {
		return CODE_M, nil
	}
	if uniformized == "F" {
		return CODE_F, nil
	}
	if uniformized == "U" {
		return CODE_U, nil
	}
	found := reCiv.FindAllString(titleDictionary.TranslateText(uniformized), 2)
	if len(found) != 1 {
		if uniformized == "" {
			return "", errors.New("unable to build a normalized string")
		}
		return CODE_U, nil
	}
	return found[0], nil
}
