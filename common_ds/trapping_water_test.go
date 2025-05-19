package common_ds

import (
	"container/list"
	"testing"
)

type Item struct {
	Val int
	Idx int
}

// calculate a relation arr: next idx with equal or greater height of current idx
func computeIdxBound(hArr []int) []int {
	stack := list.New()
	nextIdxBound := make([]int, len(hArr)) // initialize with zeroed value
	rdx := len(hArr) - 1
	stack.PushFront(Item{Val: hArr[rdx], Idx: rdx})

	for ldx := rdx - 1; ldx >= 0; ldx-- {
		lval := hArr[ldx]

		for topItem := stack.Front().Value.(Item); stack.Len() > 0 && topItem.Val < lval; {
			stack.Remove(stack.Front()) // pop every item less than current value
			topItem = stack.Front().Value.(Item)
		}
		nextIdxBound[ldx] = topItem.Idx
		stack.PushFront(Item{Val: lval, Idx: ldx})
	}

	return nextIdxBound
}

func trap(hArr []int) int {
	// find the max height

	trapped := 0
	nextIdxBound := computeIdxBound(hArr)

	for ldx := 0; ldx < len(hArr); ldx++ {

		nextIdx := nextIdxBound[ldx]

		for idx := ldx + 1; idx < nextIdx; idx++ {
			trapped += hArr[ldx] - hArr[idx]
		}
		// shift ldx to nextIdx
		ldx = nextIdx + 1
	}

	return trapped
}

func TestTrap(t *testing.T) {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	trap(arr)
}
