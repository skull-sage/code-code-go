package backtrack

func backtrackComb(curr []int, candi []int, target int, result *Result) {

	if target == 0 {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*result = append(*result, tmp)
	}

	for idx := 0; idx < len(candi); idx++ {
		c := candi[idx]

		if len(curr) > 0 && c < curr[len(curr)-1] {
			continue
		}

		if target-c >= 0 {
			curr = append(curr, candi[idx])
			backtrackComb(curr, candi, target-c, result)
			curr = curr[:len(curr)-1]
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	curr := make([]int, 0)
	result := make(Result, 0)

	backtrackComb(curr, candidates, target, &result)
	return result
}
