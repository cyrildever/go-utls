package env

import (
	"os"
	"strconv"
	"strings"
)

// GetBool tests whether the environment variable was set and eventually returns its boolean value.
func GetBool(key string) (isSet, value bool) {
	str := strings.ToLower(strings.TrimSpace(os.Getenv(key)))
	if str == "" {
		return
	}
	val, err := strconv.ParseBool(str)
	if err != nil {
		return
	}
	return true, val
}

// GetInt tests whether the environment variable was set and eventually returns its integer value.
func GetInt(key string) (isSet bool, value int) {
	str := os.Getenv(key)
	if str == "" {
		return
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		return
	}
	return true, i
}

// GetStr tests whether the environment variable was set and eventually returns its string value.
func GetStr(key string) (isSet bool, value string) {
	str := strings.TrimSpace(os.Getenv(key))
	if str == "" {
		return
	}
	return true, str
}
