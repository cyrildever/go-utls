package concurrent_test

import (
	"testing"

	"github.com/cyrildever/go-utls/common/concurrent"
	"gotest.tools/assert"
)

// TestSetAndGet ...
func TestSetAndGet(t *testing.T) {
	cmap := concurrent.NewMap()
	str := "test"

	cmap.Set("1", str)
	ret, ok := cmap.Get("1")
	assert.Equal(t, ok, true)
	if ret == nil {
		t.Fatal("Item shouldn't be nil")
	}
	assert.Equal(t, ret.(string), str)
}

// TestEmptyGet ...
func TestEmptyGet(t *testing.T) {
	cmap := concurrent.NewMap()

	ret, ok := cmap.Get("1")
	assert.Assert(t, !ok)
	if ret != nil {
		t.Fatal("Item should be nil")
	}
}

// TestMapRange ...
func TestMapRange(t *testing.T) {
	cmap := concurrent.NewMap()
	items := [3]int{1, 2, 3}
	keys := [3]string{"1", "2", "3"}
	counter := 0

	for index, item := range items {
		cmap.Set(keys[index], item)
	}

	for item := range cmap.Iter() {
		if item.Value.(int) == counter+1 {
			counter += 1
		}
	}
	assert.Equal(t, counter, 3)
}
