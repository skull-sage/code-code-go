package adhoc_dp

import (
	"fmt"
	"sort"
	"testing"
)

//type AnsArr [][]int

func buildPermute(curr []int, visitedIdx []bool, nums []int, ans *AnsArr) {
	if len(curr) == len(nums) {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*ans = append(*ans, tmp)
	}

	for i := 0; i < len(nums); i++ {
		if visitedIdx[i] || i > 0 && nums[i] == nums[i-1] && visitedIdx[i-1] {
			continue
		}

		visitedIdx[i] = true
		curr = append(curr, nums[i])
		buildPermute(curr, visitedIdx, nums, ans)
		visitedIdx[i] = false
		curr = curr[:len(curr)-1]

	}
}

func permuteUnique(nums []int) [][]int {
	sort.Sort(sort.IntSlice(nums))
	ans := make(AnsArr, 0, 0)
	curr := make([]int, 0, len(nums))
	visitIdx := make([]bool, len(nums), len(nums))
	buildPermute(curr, visitIdx, nums, &ans)
	return ans
}

func TestPermuteDuplicate(t *testing.T) {
	nums := []int{1, 1, 2}
	ans := permute(nums)
	fmt.Println(ans)
}
