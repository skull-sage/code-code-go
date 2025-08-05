package backtrack

func backtrackCombIII(start int, curr []int, candis []int, k int, target int, result *Result) {
	if target == 0 && len(curr) == k {

		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*result = append(*result, tmp)

		return
	}

	for idx := start; idx < len(candis); idx++ {
		c := candis[idx]

		if target-c >= 0 {
			curr = append(curr, c)
			backtrackCombIII(idx+1, curr, candis, k, target-c, result)
			curr = curr[:len(curr)-1]
		}
	}
}

func combinationSumIII(k int, n int) [][]int {

	candis := make([]int, 9, 9)
	for idx := range candis {
		candis[idx] = idx + 1
	}

	curr := make([]int, 0)
	result := make(Result, 0)

	backtrackCombIII(0, curr, candis, k, n, &result)
	return result
}
