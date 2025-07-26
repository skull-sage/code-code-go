package adhoc_dp

// for sorted
func twoSum(numbers []int, target int) []int {

	n := len(numbers)
	idx := 0
	jdx := n - 1
	var result []int
	for idx < n {
		delta := target - numbers[idx]

		for jdx > idx && numbers[jdx] > delta {
			jdx--
		}

		if numbers[jdx] == delta {
			result = []int{idx + 1, jdx + 1}
			break
		}

		idx++
	}

	return result
}
