package model

import (
	"reflect"
)

//--- TYPES

// Hashes is an array of Hash.
type Hashes []Hash

//--- METHODS

func (h Hashes) Len() int           { return len(h) }
func (h Hashes) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h Hashes) Less(i, j int) bool { return h[i] < h[j] }

// Equals ...
func (h *Hashes) Equals(to *Hashes) bool {
	length := h.Len()
	if length != to.Len() {
		return false
	}
	for i := 0; i < length; i++ {
		if !reflect.DeepEqual((*h)[i], (*to)[i]) {
			return false
		}
	}
	return true
}

// Contains ...
func (h Hashes) Contains(item Hash) bool {
	for _, hash := range h {
		if hash == item {
			return true
		}
	}
	return false
}
