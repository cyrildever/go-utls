package file

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// CountLines returns the number of effective lines in a file, ie. not counting any eventual last feedline/breakline
func CountLines(filepath string) (int, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := f.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			// Eventually add 1 to the count if the last byte is the line separator
			if offset, e := f.Seek(0, 2); e == nil {
				lastChar := make([]byte, 1)
				if _, e := f.ReadAt(lastChar, offset-1); e == nil {
					if !bytes.Equal(lastChar, lineSep) {
						count++
					}
				}
			}
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

// Delete ...
func Delete(filepath string) error {
	if !Exists(filepath) {
		return nil
	}
	return os.Remove(filepath)
}

// Exists ...
func Exists(filepath string) bool {
	info, err := os.Stat(filepath)
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
func GetLines(filepath string) (lines []string, err error) {
	f, err := os.Open(filepath)
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

// Truncate erases the content of a file without remove the file itself
func Truncate(filepath string, perm os.FileMode) error {
	f, err := os.OpenFile(filepath, os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return nil
}
