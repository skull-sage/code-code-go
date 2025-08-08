package adhoc_dp

func maxSubArray(nums []int) int {
	max := nums[0]
	sum_track := make([]int, len(nums))
	idx_track := make([]int, len(nums))

	sum_track[0] = nums[0]
	idx_track[0] = 0

	for idx := 1; idx < len(nums); idx++ {
		c := nums[idx]

		if sum_track[idx-1]+c > c {
			sum_track[idx] = sum_track[idx-1] + c
			idx_track[idx] = idx_track[idx-1]
		} else {
			sum_track[idx] = c
			idx_track[idx] = idx
		}

		if sum_track[idx] > max {
			max = sum_track[idx]
		}

	}
	return max
}
