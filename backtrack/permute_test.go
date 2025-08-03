package adhoc_dp

import (
	"sort"
)

type AnsArr [][]int

func recurPermute(idx int, nums []int, ans *AnsArr) {

	if idx == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*ans = append(*ans, tmp)
	}

	for i := idx; i < len(nums); i++ {
		if i > idx && nums[i] == nums[idx] {
			continue
		}
		nums[idx], nums[i] = nums[i], nums[idx]
		recurPermute(idx+1, nums, ans)
		nums[idx], nums[i] = nums[i], nums[idx]
	}

}

func permute(nums []int) [][]int {

	sort.Sort(sort.IntSlice(nums))

	ans := make(AnsArr, 0, 0)

	recurPermute(0, nums, &ans)
	return ans
}
