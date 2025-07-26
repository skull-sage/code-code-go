package adhoc_dp

import "testing"

func twoSum(nums []int, target int) []int {
	idxMap := make(map[int]int)

	for idx, num := range nums {
		idxMap[num] = idx
	}

	var result []int
	for idx, num := range nums {
		delta := target - num
		deltaIdx, ok := idxMap[delta]

		// in case of two equalk value map will contain right most index
		// hence, index of delta should be > idx for equal val: target = 6, num = 3 then delta is also 3
		if deltaIdx != idx && ok {
			result = []int{idx, deltaIdx}
			break
		}
	}

	return result

}

func TestTwoSum(t *testing.T) {

}
