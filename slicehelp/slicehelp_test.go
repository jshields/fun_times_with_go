package slicehelp

import "testing"

func TestIntSlicesEqual(t *testing.T) {
	testCases := map[string]struct {
		inputArr1 []int
		inputArr2 []int
		expected  bool
	}{
		"nil arr1":                        {inputArr1: nil, inputArr2: []int{1}, expected: false},
		"empty arr1":                      {inputArr1: []int{}, inputArr2: []int{1}, expected: false},
		"different lengths":               {inputArr1: []int{1, 2}, inputArr2: []int{1, 2, 3}, expected: false},
		"same length, different contents": {inputArr1: []int{1, 2}, inputArr2: []int{2, 3}, expected: false},
		"both nil":                        {inputArr1: nil, inputArr2: nil, expected: true},
		"both empty":                      {inputArr1: []int{}, inputArr2: []int{}, expected: true},
		"same contents":                   {inputArr1: []int{1, 2, 3}, inputArr2: []int{1, 2, 3}, expected: true},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := IntSlicesEqual(tc.inputArr1, tc.inputArr2)
			if actual != tc.expected {
				t.Fatalf("actual %v != expected %v", actual, tc.expected)
			}
		})
	}
}
