package sort

func bubble(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		// -i because i-th iteration has bubbled up the i-th largest number
		// -1 because we look ahead by 1 item
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// no need for tmp variable when using multiple assignment
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
