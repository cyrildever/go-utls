package utils

import (
	"reflect"
)

// IsPointer returns `true` if the passed item is a pointer, not a value
func IsPointer(item interface{}) bool {
	return reflect.ValueOf(item).Kind() == reflect.Ptr
}

// IsValue returns `true` if the passed item is a value, not a pointer
func IsValue(item interface{}) bool {
	return !IsPointer(item)
}
