package concurrent

import (
	"errors"
	"math/rand"
	"sync"

	cerrors "github.com/cyrildever/go-utls/common/concurrent/errors"
)

//--- TYPES

// Slice type that can be safely shared between goroutines
type Slice struct {
	sync.RWMutex
	items         []interface{}
	hasInitLength bool
	lastIndex     int
	length        int
}

// SliceItem is a concurrent slice item
type SliceItem struct {
	Index int
	Value interface{}
}

// Checker is a function type. You can see it as a function that you have to implement for each structure you want to add inside the slice.
// It should be passed as parameter at the Check() or Get() methods.
type Checker func(interface{}, interface{}) bool

//--- METHODS

// Append adds an item to the concurrent slice
func (cs *Slice) Append(item interface{}) {
	cs.Lock()
	defer cs.Unlock()

	if !cs.hasInitLength {
		cs.items = append(cs.items, item)
	} else {
		if cs.lastIndex < cs.length {
			cs.items[cs.lastIndex] = item
			cs.lastIndex++
		} else {
			cs.items = append(cs.items, item)
			cs.hasInitLength = false
		}
	}
}

// Get will return the interface corresponding with the value and the Checker given.
// If no Checker is given, then we're comparing interface.
func (cs *Slice) Get(value interface{}, f Checker) interface{} {
	cs.RLock()
	defer cs.RUnlock()

	if f != nil {
		for _, item := range cs.items {
			if f(item, value) {
				return item
			}
		}
	} else {
		for _, item := range cs.items {
			if item == value {
				return item
			}
		}
	}
	return nil
}

// Check tests if the passed value exists in the concurrent slice
func (cs *Slice) Check(value interface{}, f Checker) bool {
	cs.RLock()
	defer cs.RUnlock()

	if f != nil {
		for _, item := range cs.items {
			if f(item, value) {
				return true
			}
		}
	} else {
		for _, item := range cs.items {
			if item == value {
				return true
			}
		}
	}
	return false
}

// Delete removes the specified item from the slice
func (cs *Slice) Delete(value interface{}, f Checker) error {
	cs.Lock()
	defer cs.Unlock()

	if cs.size() == 0 {
		return cerrors.EmptySlice{}
	}

	if f != nil {
		for i, item := range cs.items {
			if f(item, value) {
				cs.deleteItem(i)
				return nil
			}
		}
	} else {
		for i, item := range cs.items {
			if item == value {
				cs.deleteItem(i)
				return nil
			}
		}
	}

	return errors.New("impossible to find item")
}

// See https://yourbasic.org/golang/delete-element-slice/
// Order isn't important for us; speed is.
func (cs *Slice) deleteItem(index int) {
	if cs.size() < 2 {
		cs.clear()
		return
	}
	last := len(cs.items) - 1
	cs.items[index] = cs.items[last]
	cs.items[last] = nil
	cs.items = cs.items[:last]
}

// GetAll returns all items in the concurrent slice present when the caller calls it
func (cs *Slice) GetAll() ([]interface{}, bool) {
	cs.RLock()
	defer cs.RUnlock()

	if cs.size() > 0 {
		return cs.items, true
	}

	return nil, false
}

// GetAt returns the item at the passed index
//
// NB: It returns nil if the index is out of bound or there is no item
func (cs *Slice) GetAt(index int) interface{} {
	cs.RLock()
	defer cs.RUnlock()

	if index < 0 || index >= len(cs.items) {
		return nil
	}

	return cs.items[index]
}

// GetOne returns one random item
func (cs *Slice) GetOne() (interface{}, bool) {
	cs.RLock()
	defer cs.RUnlock()

	if cs.size() == 0 {
		return nil, false
	}

	randIndex := rand.Intn(cs.size())
	return cs.items[randIndex], true
}

// Iter iterates over the items in the concurrent slice
// Each item is sent over a channel, so that
// we can iterate over the slice using the built-in range keyword
func (cs *Slice) Iter() <-chan SliceItem {
	cs.RLock()
	defer cs.RUnlock()

	c := make(chan SliceItem)

	f := func() {
		for index, value := range cs.items {
			c <- SliceItem{index, value}
		}
		close(c)
	}
	go f()

	return c
}

// Pop returns the last item in the concurrent slice and remove it for the latter
func (cs *Slice) Pop() (interface{}, error) {
	cs.Lock()
	defer cs.Unlock()

	csLen := len((*cs).items)
	if csLen > 0 {
		item := (*cs).items[csLen-1]
		(*cs).items = (*cs).items[:csLen-1]
		return item, nil
	}
	return nil, cerrors.EmptySlice{}
}

// PopAll returns all the items of the concurrent slice and empties it
func (cs *Slice) PopAll() ([]interface{}, error) {
	cs.Lock()
	defer cs.Unlock()

	csLen := len((*cs).items)
	if csLen > 0 {
		items := (*cs).items[:]
		cs.clear()
		return items, nil
	}
	return nil, cerrors.EmptySlice{}
}

// Size returns the length of the concurrent slice
func (cs *Slice) Size() int {
	cs.RLock()
	defer cs.RUnlock()

	return cs.size()
}

func (cs *Slice) size() int {
	return len(cs.items)
}

func (cs *Slice) clear() {
	(*cs).items = []interface{}{}
}

//--- FUNCTIONS

// NewSlice ...
func NewSlice(size ...int) *Slice {
	var items []interface{}
	hasInitLength := false
	length := 0
	if len(size) > 0 && size[0] > 0 {
		length = size[0]
		items = make([]interface{}, length)
		hasInitLength = true
	}
	return &Slice{
		items:         items,
		hasInitLength: hasInitLength,
		length:        length,
	}
}
