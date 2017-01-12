package goutils

// SliceContains returns true if a slice of strings includes a specific string
func SliceContains(needle string, haystack []string) bool {
	for _, value := range haystack {
		if needle == value {
			return true
		}
	}

	return false
}
