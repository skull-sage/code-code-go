package adhoc_dp

func countBits(n int) []int {

	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}

	bitCount := make([]int, n+1, n+1)
	bitCount[0] = 0
	bitCount[1] = 1

	for i := 2; i <= n; i++ {
		j := i & (i - 1)
		bitCount[i] = bitCount[j] + 1
	}

	return bitCount
}
