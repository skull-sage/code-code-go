package adhoc_dp

func longestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	existMap := make(map[int]int)

	for _, num := range nums {
		existMap[num] = 1
	}

	maxCount := 1
	for key, _ := range existMap {
		count := 0
		for val := key; existMap[val] == 1; val++ {
			existMap[val] = 0 // as we already have counted
			count++
		}

		for val := key - 1; existMap[val] == 1; val-- {
			existMap[val] = 0
			count++
		}

		if maxCount < count {
			maxCount = count
		}
	}

	return maxCount

}
