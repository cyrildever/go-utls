package utils

// ContainString returns true if the passed slice of strings contains the passed string value
func ContainString(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
