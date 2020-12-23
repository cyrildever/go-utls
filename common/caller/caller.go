package caller

import (
	"fmt"
	"runtime"
)

// Get returns the calling location (file and line number) if found, eg.
//	if file, line, ok := caller.Get(); ok {
//		// /path/to/calling_file.go  11
//		fmt.Println(file, line)
//	}
func Get() (file string, line int, found bool) {
	if _, f, l, ok := runtime.Caller(2); ok {
		return f, l, true
	}
	return
}

// AsString returns a pre-formatted string from the caller.Get() function using the following format: `fileName#lineNumber`, eg.
//	// /path/to/calling_file.go#22
//	fmt.Println(caller.AsString())
func AsString() string {
	if _, f, l, ok := runtime.Caller(2); ok {
		return fmt.Sprintf("%s#%d", f, l)
	}
	return ""
}
