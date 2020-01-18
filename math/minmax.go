package math

// Min finds the smallest integer from the arguments.
// Returns zero if no arguments are passed.
func Min(nums ...int) int {
	var min int
	if len(nums) == 0 {
		return 0
	}
	// min must start at one of the numbers in the collection,
	// since the collection does not necessarily contain the zero value of an iteger type (0).
	min = nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min
}

// Max finds the largest integer from the arguments.
// Returns zero if no arguments are passed.
func Max(nums ...int) int {
	var max int
	if len(nums) == 0 {
		return 0
	}
	max = nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

// MinMax finds the pair that is the smallest and largest integer from the arguments.
// Returns zeroes if no arguments are passed.
func MinMax(nums ...int) (min int, max int) {
	// if only a single number was passed, it is the min and max
	if len(nums) == 0 {
		return
	}
	// start min and max at valid values from the collection, not zeroes
	min = nums[0]
	max = nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return
}
