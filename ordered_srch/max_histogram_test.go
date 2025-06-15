package ordered_srch

import (
	"fmt"
	"sort"
	"testing"
)

type element struct {
	val int
	idx int
}

func toElementArr(arr []int) []element {

	elmArr := make([]element, 0, len(arr))

	for i := 0; i < len(arr); i++ {
		elmArr = append(elmArr, element{arr[i], i})
	}

	return elmArr
}

func maxHistogram(arr []int) int {
	maxArea := 0
	elmArr := toElementArr(arr)

	// sort.Slice works in-place
	sort.Slice(elmArr, func(i, j int) bool {
		if elmArr[i].val == elmArr[j].val {
			return elmArr[i].idx > elmArr[j].idx
		}
		return elmArr[i].val > elmArr[j].val

	})

	fmt.Println(elmArr)

	phi := make([]int, len(elmArr), len(elmArr))

	for i := 0; i < len(elmArr); i++ {
		val := elmArr[i]
	}

	// we start with currMax,
	// we go left until ldx-th element < currMax
	// we go right until rdx-th element < currMax
	// we calculate area = currMax * (rdx - ldx)
	// if we meet an idx that is already covered we take (rdx - ldx)

	return maxArea

}

func TestHisto(t *testing.T) {
	arr := []int{2, 6, 6, 1, 2, 2, 5, 6, 2, 6, 3}
	result := maxHistogram(arr)

	fmt.Println(result)

}
