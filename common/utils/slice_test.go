package utils_test

import (
	"testing"

	"github.com/cyrildever/go-utls/common/utils"
	"gotest.tools/assert"
)

// TestChunk ...
func TestChunk(t *testing.T) {
	slice := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	chunks := utils.Chunk(slice, 3)
	assert.Equal(t, len(chunks), 3)
	assert.DeepEqual(t, chunks[0], []string{"1", "2", "3"})
	assert.DeepEqual(t, chunks[1], []string{"4", "5", "6"})
	assert.DeepEqual(t, chunks[2], []string{"7", "8"})

	chunks = utils.Chunk(slice, 9)
	assert.Equal(t, len(chunks), 1)
	assert.DeepEqual(t, chunks[0], slice)
}

// TestContainString ...
func TestContainString(t *testing.T) {
	slice := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	found := utils.ContainString(slice, "1")
	assert.Assert(t, found)

	slice = make([]string, 4)
	found = utils.ContainString(slice, "1")
	assert.Assert(t, found == false)
}
