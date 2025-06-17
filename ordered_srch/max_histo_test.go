package ordered_srch

import (
	"fmt"
	"testing"
)

type element struct {
	val   int
	start int
	end   int
}

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0, 0)}
}

func (s *Stack[T]) IsEmpty() bool {
	return len((*s).items) == 0
}

func (s *Stack[T]) Push(elm T) {
	s.items = append(s.items, elm)
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	lastIdx := len(s.items) - 1
	item := s.items[lastIdx]
	s.items = s.items[:lastIdx]
	return item, true

}

func (s *Stack[T]) Clear() {
	s.items = s.items[:0]
}

func maxHistogram(arr []int) int {

	stack := NewStack[element]()
	//appending zero for all pops
	arr = append(arr, 0)
	maxArea := 0

	fmt.Println("input: ", arr)

	for i := range arr {
		val := arr[i]

		start := i
		currTop, hasTop := stack.Peek()
		topStart := currTop.start
		for hasTop && currTop.val >= val {
			area := currTop.val * (topStart - currTop.start + 1)
			if maxArea < area {
				maxArea = area
			}
			start = currTop.start

			stack.Pop()
			currTop, hasTop = stack.Peek()
		}

		stack.Push(element{val, start, i})
		//fmt.Println("# area", maxArea)
		//	fmt.Println("# stack", stack.items)
		//	fmt.Println()
	}

	return maxArea
}

func TestHisto(t *testing.T) {
	arr := []int{1, 4, 4, 5, 4, 6, 7, 3}
	result := maxHistogram(arr)

	fmt.Println(result)

}
