package file_test

import (
	"io/ioutil"
	"testing"

	"github.com/cyrildever/go-utls/common/file"
	"gotest.tools/assert"
)

// TestFile ...
func TestFile(t *testing.T) {
	testfile := "testFile.txt"
	content := "Ceci est un test"
	err := ioutil.WriteFile(testfile, []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}

	assert.Assert(t, file.Exists(testfile))
	lines, err := file.GetLines(testfile)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(lines), 1)
	assert.Equal(t, lines[0], content)

	err = file.Truncate(testfile, 0644)
	if err != nil {
		t.Fatal(err)
	}
	assert.Assert(t, file.Exists(testfile))
	lines, err = file.GetLines(testfile)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(lines), 0)

	err = file.Delete(testfile)
	assert.Assert(t, err == nil)
}
