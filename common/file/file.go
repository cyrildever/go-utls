package file

import (
	"bufio"
	"os"
)

// Exists ...
func Exists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Find recursively searches for a valid file path in a file tree
func Find(filepath string) (string, bool) {
	MAX_ITERATION := 16 // Adapt if max depth of folder goes beyond
	i := 0
	found := Exists(filepath)
	for !found && i < MAX_ITERATION {
		filepath = "../" + filepath
		found = Exists(filepath)
		i++
	}
	if found {
		return filepath, true
	}
	return "", false
}

// GetLines ...
func GetLines(filename string) (lines []string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return
}
