package errors

// EmptySlice ...
type EmptySlice struct{}

// Error ...
func (err EmptySlice) Error() (r string) {
	r += "slice is empty"
	return
}
