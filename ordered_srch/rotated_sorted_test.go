package ordered_srch

func findMinIdx(nums []int) int {
	minIdx := nums[0]

	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] < nums[minIdx] {
			minIdx = idx
		}
	}
	return minIdx
}

func binarySearch[T any](arr []T, x T) int {
	
}

func search(nums []int, target int) int {
	minIdx := findMinIdx(nums)
	if nums[minIdx] == target {
		return minIdx
	}
}
