package concurrent

import (
	"sync"

	"github.com/elliotchance/orderedmap"
)

// Map type that can be safely shared between
// goroutines that require read/write access to a map
type Map struct {
	sync.RWMutex
	items *orderedmap.OrderedMap
}

// NewMap ...
func NewMap() *Map {
	cm := &Map{
		items: orderedmap.NewOrderedMap(),
	}
	return cm
}

// MapItem represents a concurrent map item
type MapItem struct {
	Key   string
	Value interface{}
}

// Set a key in a concurrent map
func (cm *Map) Set(key string, value interface{}) bool {
	cm.Lock()
	defer cm.Unlock()

	return cm.items.Set(key, value)
}

// Get a key from a concurrent map
func (cm *Map) Get(key string) (interface{}, bool) {
	cm.RLock()
	defer cm.RUnlock()

	return cm.items.Get(key)
}

// GetAll returns all the items of the concurrent map
func (cm *Map) GetAll() ([]interface{}, bool) {
	if cm.Size() > 0 {
		var items []interface{}
		for item := range cm.Iter() {
			items = append(items, item.Value)
		}
		return items, true
	}
	return nil, false
}

// Delete a key from a concurrent map
func (cm *Map) Delete(key string) bool {
	cm.Lock()
	defer cm.Unlock()

	return cm.items.Delete(key)
}

// Keys returns all the keys of the map
func (cm *Map) Keys() (keys []string) {
	if cm.Size() > 0 {
		keys = make([]string, cm.Size())
		i := 0
		for k := range cm.Iter() {
			keys[i] = k.Key
			i++
		}
	}
	return
}

// Pop a key from a concurrent map
func (cm *Map) Pop(key string) (interface{}, bool) {
	cm.Lock()
	defer cm.Unlock()

	value, ok := cm.items.Get(key)
	if ok {
		cm.items.Delete(key)
	}

	return value, ok
}

// Iter iterates over the items in a concurrent map.
// Each item is sent over a channel, so that
// we can iterate over the map using the builtin range keyword.
func (cm *Map) Iter() <-chan MapItem {
	c := make(chan MapItem)

	f := func() {
		cm.RLock()
		defer cm.RUnlock()

		for el := cm.items.Front(); el != nil; el = el.Next() {
			c <- MapItem{el.Key.(string), el.Value}
		}
		close(c)
	}
	go f()

	return c
}

// Size ...
func (cm *Map) Size() int {
	cm.RLock()
	defer cm.RUnlock()

	return cm.size()
}

func (cm *Map) size() int {
	return cm.items.Len()
}
