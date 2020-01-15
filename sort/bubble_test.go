package sort

import "testing"

func slicesEqual(arr1, arr2 []int) bool {
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

func Test_bubble(t *testing.T) {
	testCases := map[string]struct {
		input    []int
		expected []int
	}{
		"empty":                  {input: []int{}, expected: []int{}},
		"zero":                   {input: []int{0}, expected: []int{0}},
		"3-2-1":                  {input: []int{3, 2, 1}, expected: []int{1, 2, 3}},
		"with duplicates":        {input: []int{2, 1, 8, 5, 5, 2, 3, 2, 3, 7}, expected: []int{1, 2, 2, 2, 3, 3, 5, 5, 7, 8}},
		"some big numbers":       {input: []int{500, 200, 1000, 2, 70, 99999}, expected: []int{2, 70, 200, 500, 1000, 99999}},
		"negative numbers only":  {input: []int{-5, -2, -6, -1}, expected: []int{-6, -5, -2, -1}},
		"negative numbers mixed": {input: []int{1, -1, 2, -3, 5, 9, -9}, expected: []int{-9, -3, -1, 1, 2, 5, 9}},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := bubble(tc.input)
			if !slicesEqual(actual, tc.expected) {
				t.Fatalf("%s: actual %v != expected %v", name, actual, tc.expected)
			}
		})
	}
}
