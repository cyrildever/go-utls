package model

//--- TYPES

// Keys is an array of Key.
type Keys []Key

//--- METHODS

// Contains ...
func (k Keys) Contains(item Key) bool {
	for _, key := range k {
		if key == item {
			return true
		}
	}
	return false
}
