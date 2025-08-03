package adhoc_dp

import (
	"sort"
	"testing"
)

type ResultArr [][]int

func recurSubset(idx int, curr []int, nums []int, result *ResultArr) {

	tmp := make([]int, len(curr))
	copy(tmp, curr)
	*result = append(*result, tmp)

	l := len(curr)

	for i := idx + 1; i < len(nums); i++ {

		if l > 0 && nums[i] <= curr[l-1] {
			continue
		}

		curr = append(curr, nums[i]) // {1} / {2} / {3} / {4}
		//fmt.Println("idx:", idx, "i:", i, curr)
		recurSubset(idx+1, curr, nums, result)
		curr = curr[:l]
	}

}

func subsets(nums []int) [][]int {
	sort.Ints(nums)
	ans := make(ResultArr, 0)
	curr := make([]int, 0)
	recurSubset(-1, curr, nums, &ans)
	return ans
}

func TestSubset(t *testing.T) {

	nums := []int{1, 2, 3, 4}
	ans := subsets(nums)
	t.Log(ans)
}
