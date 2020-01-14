package slicehelp

// IntSlicesEqual compares equality of two slices of ints.
func IntSlicesEqual(arr1, arr2 []int) bool {
	if (arr1 == nil) != (arr2 == nil) {
		return false
	}
	if len(arr1) != len(arr2) {
		return false
	}
	for ind, val := range arr1 {
		if val != arr2[ind] {
			return false
		}
	}
	return true
}
