package backtrack

import (
	"fmt"
	"sort"
	"testing"
)

type ResultArr [][]int

func recurSubset(idx int, curr []int, nums []int, result *ResultArr) {

	tmp := make([]int, len(curr))
	copy(tmp, curr)
	*result = append(*result, tmp)

	l := len(curr)

	for i := idx; i < len(nums); i++ {

		curr = append(curr, nums[i]) // {1} / {2} / {3} / {4}
		//fmt.Println("idx:", idx, "i:", i, curr)
		recurSubset(i+1, curr, nums, result)
		curr = curr[:l]
	}

}

func subsets(nums []int) [][]int {
	sort.Ints(nums)
	ans := make(ResultArr, 0)
	curr := make([]int, 0)
	recurSubset(0, curr, nums, &ans)
	return ans
}

func TestSubset(t *testing.T) {

	nums := []int{1, 2, 3, 4}
	ans := subsets(nums)
	for _, a := range ans {
		fmt.Println(a)
	}
}
