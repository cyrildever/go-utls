package model

import (
	"reflect"
)

//--- TYPES

// Keys is an array of Key.
type Keys []Key

//--- METHODS

func (k Keys) Len() int           { return len(k) }
func (k Keys) Swap(i, j int)      { k[i], k[j] = k[j], k[i] }
func (k Keys) Less(i, j int) bool { return k[i] < k[j] }

// Equals ...
func (k *Keys) Equals(to *Keys) bool {
	length := k.Len()
	if length != to.Len() {
		return false
	}
	for i := 0; i < length; i++ {
		if !reflect.DeepEqual((*k)[i], (*to)[i]) {
			return false
		}
	}
	return true
}

// Contains ...
func (k Keys) Contains(item Key) bool {
	for _, key := range k {
		if key == item {
			return true
		}
	}
	return false
}
