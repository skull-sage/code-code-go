package adhoc_dp

func productExceptSelf(nums []int) []int {
	n := len(nums)
	prefixProd := make([]int, n, n)
	suffixProd := make([]int, n, n)

	prefixProd[0] = nums[0]
	for idx := 1; idx < n; idx++ {
		prefixProd[idx] = prefixProd[idx-1] * nums[idx]
	}

	suffixProd[n-1] = nums[n-1]
	for idx := n - 2; idx >= 0; idx-- {
		suffixProd[idx] = suffixProd[idx+1] * nums[idx]
	}

	result := make([]int, n, n)
	result[0] = suffixProd[1]
	result[n-1] = prefixProd[n-2]

	for idx := 1; idx < n-1; idx++ {
		result[idx] = prefixProd[idx-1] * suffixProd[idx+1]
	}

	return result

}
