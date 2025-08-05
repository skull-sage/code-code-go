package backtrack

import "sort"

type Result [][]int

func backtrackCombII(start int, curr []int, candi []int, target int, result *Result) {

	if target == 0 {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*result = append(*result, tmp)
	}

	prevCandi := -1
	for idx := start; idx < len(candi); idx++ {
		c := candi[idx]

		if prevCandi == c {
			continue
		} else {
			prevCandi = c
		}

		if target-c >= 0 {
			curr = append(curr, candi[idx])
			backtrackCombII(idx+1, curr, candi, target-c, result)
			curr = curr[:len(curr)-1]
		}
	}
}

func combinationSumII(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	curr := make([]int, 0)
	result := make(Result, 0)

	backtrackCombII(0, curr, candidates, target, &result)
	return result
}
