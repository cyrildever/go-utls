package env

import (
	"os"
	"strconv"
	"strings"
)

// GetBool tests whether the environment variable was set and eventually returns its boolean value or a default if passed.
// In the last case, the environment variable itself would be set.
func GetBool(key string, defaultValue ...bool) (value, wasSet bool) {
	str := strings.ToLower(strings.TrimSpace(os.Getenv(key)))
	if str == "" {
		if len(defaultValue) == 1 {
			os.Setenv(key, strconv.FormatBool(defaultValue[0]))
			return defaultValue[0], false
		}
		return
	}
	val, err := strconv.ParseBool(str)
	if err != nil {
		return
	}
	return val, true
}

// GetInt tests whether the environment variable was set and eventually returns its integer value or a default if passed.
// In the last case, the environment variable itself would be set.
func GetInt(key string, defaultValue ...int) (value int, wasSet bool) {
	str := os.Getenv(key)
	if str == "" {
		if len(defaultValue) == 1 {
			os.Setenv(key, strconv.FormatInt(int64(defaultValue[0]), 10))
			return defaultValue[0], false
		}
		return
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		return
	}
	return i, true
}

// GetStr tests whether the environment variable was set and eventually returns its string value.
// In the last case, the environment variable itself would be set.
func GetStr(key string, defaultValue ...string) (value string, wasSet bool) {
	str := strings.TrimSpace(os.Getenv(key))
	if str == "" {
		if len(defaultValue) == 1 {
			os.Setenv(key, defaultValue[0])
			return defaultValue[0], false
		}
		return
	}
	return str, true
}
