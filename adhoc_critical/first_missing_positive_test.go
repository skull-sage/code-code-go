package adhoc_critical

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v

}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for idx, val := range nums {
		if val >= n+1 || val <= 0 {
			nums[idx] = n + 1
		}
	}
	// fmt.Println(nums)
	for idx := 0; idx < n; idx++ {
		val := abs(nums[idx])
		if val < n+1 && nums[val-1] > 0 {
			nums[val-1] = -nums[val-1]
		}
	}
	//fmt.Println(nums)

	for idx := 1; idx <= n; idx++ {
		if nums[idx-1] < 0 == false {
			return idx
		}
	}

	return n + 1
}
