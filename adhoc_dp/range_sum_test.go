package adhoc_dp

import (
	"fmt"
	"testing"
)

type NumArray struct {
	prefixSum []int
	numArr    []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	prefixSum := make([]int, n, n)
	prefixSum[0] = nums[0]
	for idx := 1; idx < n; idx++ {
		prefixSum[idx] = prefixSum[idx-1] + nums[idx]
	}

	return NumArray{prefixSum, nums}
}

func (this *NumArray) SumRange(left int, right int) int {

	if left == 0 {
		return this.prefixSum[right]
	}

	return this.prefixSum[right] - this.prefixSum[left] + this.numArr[left]
}

func TestRangeSum(t *testing.T) {
	sumCont := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(sumCont.prefixSum)
	fmt.Println(sumCont.SumRange(0, 2))
	fmt.Println(sumCont.SumRange(2, 5))
	fmt.Println(sumCont.SumRange(0, 5))

}
