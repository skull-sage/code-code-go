package backtrack

import (
	"fmt"
	"sort"
	"testing"
)

func backtrack(start int, curr []int, nums []int, result *ResultArr) {

	tmp := make([]int, len(curr))
	copy(tmp, curr)
	*result = append(*result, tmp)

	l := len(curr)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}

		curr = append(curr, nums[i]) // {1} / {2} / {3} / {4}
		//fmt.Println("start:", start, "i:", i, curr)
		backtrack(i+1, curr, nums, result)
		curr = curr[:l]
	}

}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ans := make(ResultArr, 0)
	curr := make([]int, 0)
	backtrack(0, curr, nums, &ans)
	return ans
}

func TestSubsetII(t *testing.T) {

	nums := []int{1, 2, 2, 3}
	ans := subsetsWithDup(nums)
	//t.Log(ans)
	for _, a := range ans {
		fmt.Println(a)
	}
}
