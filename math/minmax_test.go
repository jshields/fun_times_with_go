package math

import "testing"

func Test_Min(t *testing.T) {
	testCases := map[string]struct {
		input    []int
		expected int
	}{
		"empty": {
			expected: 0,
		},
		"zero": {
			input:    []int{0},
			expected: 0,
		},
		"all the same positive": {
			// (this would break if min starts at zero when zero is not actually in the collection)
			input:    []int{3, 3, 3},
			expected: 3,
		},
		"all the same negative": {
			input:    []int{-2, -2, -2},
			expected: -2,
		},
		"negative numbers": {
			input:    []int{-3, -1, -5, -8, -4},
			expected: -8,
		},
		"negative and positive numbers": {
			input:    []int{1, -3, 2, 5, -1},
			expected: -3,
		},
		"positive numbers": {
			input:    []int{4, 2, 6, 11, 9},
			expected: 2,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := Min(tc.input...)
			if actual != tc.expected {
				t.Fatalf("%s: actual %v != expected %v", name, actual, tc.expected)
			}
		})
	}
}

func Test_Max(t *testing.T) {
	testCases := map[string]struct {
		input    []int
		expected int
	}{
		"empty": {
			expected: 0,
		},
		"zero": {
			input:    []int{0},
			expected: 0,
		},
		"all the same positive": {
			input:    []int{3, 3, 3},
			expected: 3,
		},
		"all the same negative": {
			input:    []int{-2, -2, -2},
			expected: -2,
		},
		"negative numbers": {
			input:    []int{-3, -1, -5, -8, -4},
			expected: -1,
		},
		"negative and positive numbers": {
			input:    []int{1, -3, 2, 5, -1},
			expected: 5,
		},
		"positive numbers": {
			input:    []int{4, 2, 6, 11, 9},
			expected: 11,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := Max(tc.input...)
			if actual != tc.expected {
				t.Fatalf("%s: actual %v != expected %v", name, actual, tc.expected)
			}
		})
	}
}
func Test_MinMax(t *testing.T) {
	testCases := map[string]struct {
		input       []int
		expectedMin int
		expectedMax int
	}{
		"empty": {
			expectedMin: 0,
			expectedMax: 0,
		},
		"zero": {
			input:       []int{0},
			expectedMin: 0,
			expectedMax: 0,
		},
		"all the same positive": {
			input:       []int{3, 3, 3},
			expectedMin: 3,
			expectedMax: 3,
		},
		"all the same negative": {
			input:       []int{-2, -2, -2},
			expectedMin: -2,
			expectedMax: -2,
		},
		"negative numbers": {
			input:       []int{-3, -1, -5, -8, -4},
			expectedMin: -8,
			expectedMax: -1,
		},
		"negative and positive numbers": {
			input:       []int{1, -3, 2, 5, -1},
			expectedMin: -3,
			expectedMax: 5,
		},
		"positive numbers": {
			input:       []int{4, 2, 6, 11, 9},
			expectedMin: 2,
			expectedMax: 11,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actualMin, actualMax := MinMax(tc.input...)
			if actualMin != tc.expectedMin {
				t.Fatalf("%s: actualMin %v != expectedMin %v", name, actualMin, tc.expectedMin)
			}
			if actualMax != tc.expectedMax {
				t.Fatalf("%s: actualMax %v != expectedMax %v", name, actualMax, tc.expectedMax)
			}
		})
	}
}
