package ordered_srch

import (
	"fmt"
	"testing"
)

type element struct {
	val        int
	expandLeft int
}

type Stack[T any] struct {
	items []T
}

func NewStack() {
	
}
func (s *Stack[T]) push(elm T) {
	s.items = append(items, elm)
}

func (s *Stack[T]) pop(elm T) {

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
	//elmArr := toElementArr(arr)

	for i := 0; i < len(arr); i++ {

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
