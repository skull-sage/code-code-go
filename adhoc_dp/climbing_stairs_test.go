package adhoc_dp

func climbStairs(n int) int {

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	climbMemo := make([]int, n+1, n+1)
	climbMemo[0] = 0
	climbMemo[1] = 1
	climbMemo[2] = 2

	if n >= 3 {
		for idx := 3; idx <= n; idx++ {
			climbMemo[idx] = climbMemo[idx-1] + climbMemo[idx-2]
		}
	}

	return climbMemo[n]

}
