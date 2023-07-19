package check

// InSlice - check if element is in slice
func InSlice(slice []string, elem string) bool {

	for _, str := range slice {

		if str == elem {
			return true
		}
	}

	return false
}
