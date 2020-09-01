package concurrent_test

import (
	"testing"

	"github.com/cyrildever/go-utls/common/concurrent"
	"github.com/cyrildever/go-utls/common/concurrent/errors"
	"gotest.tools/assert"
)

// TestAddAndPop ...
func TestAddAndPop(t *testing.T) {
	slice := concurrent.Slice{}
	str := "test"

	slice.Append(str)
	ret, err := slice.Pop()
	if err != nil {
		t.Errorf("no error expected, but error: %v", err)
	}
	if ret == nil {
		t.Errorf("item shouldn't be nil")
	}
	assert.Equal(t, ret.(string), str, "item is not the same")
}

// TestEmptyPop ...
func TestEmptyPop(t *testing.T) {
	slice := concurrent.Slice{}

	ret, err := slice.Pop()
	if err == nil {
		t.Errorf("error was expected but got nil")
	}
	assert.Equal(t, ret, nil, "returned item should be nil")
	assert.Equal(t, err, errors.EmptySlice{}, "error is not in the good format")
}

// TestRange ...
func TestSliceRange(t *testing.T) {
	slice := concurrent.Slice{}
	items := [3]int{1, 2, 3}
	counter := 0

	for _, item := range items {
		slice.Append(item)
	}
	for item := range slice.Iter() {
		if item.Value.(int) == counter+1 {
			counter += 1
		}
	}
	assert.Equal(t, counter, 3, "cannot found all the appended items")
}

// TestPopAll ...
func TestPopAll(t *testing.T) {
	slice := concurrent.Slice{}
	str1 := "test1"
	slice.Append(str1)
	str2 := "test2"
	slice.Append(str2)

	// Before
	assert.Equal(t, slice.Size(), 2, "not the right number of items")

	// After
	items, _ := slice.PopAll()
	assert.Equal(t, len(items), 2, "not the good length for the returned array of items")
	assert.Equal(t, slice.Size(), 0, "the slice should be empty after invoking PopAll()")
}
