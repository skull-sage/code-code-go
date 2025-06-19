package ordered_srch

func findMinIdx(nums []int) int {
	minIdx := 0

	for idx := 1; idx < len(nums); idx++ {
		if nums[idx] < nums[minIdx] {
			minIdx = idx
		}
	}
	return minIdx
}

// generic form [T constraints.Ordered] to support
// comparing ops(>, <, >=, <=)
// but in leetcode not possible to import "golang.org/x/exp/constraints"
func binarySearch(arr []int, x int) (int, bool) {
	n := len(arr)
	idx := -1
	found := false
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == x {
			idx, found = mid, true
			break
		} else if arr[mid] < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return idx, found
}

func search(nums []int, target int) int {
	minIdx := findMinIdx(nums)
	if nums[minIdx] == target {
		return minIdx
	}

	idx, found := binarySearch(nums[:minIdx], target)
	if found {
		return idx
	}
	idx, found = binarySearch(nums[minIdx:], target)
	if found {
		return idx + minIdx
	}
	return -1
}
