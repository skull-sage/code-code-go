package common_ds

import (
	"fmt"
	"testing"
)

// calculate a relation arr: next idx with equal or greater height of current idx
func computeIdxBound(hArr []int) []int {
	size := len(hArr)

	nextHigherIdx := make([]int, size, size) // initialize with zeroed value
	stack := make([]int, 0, size)
	for ldx := size - 1; ldx >= 0; ldx-- {
		lval := hArr[ldx]
		nextHigh := ldx
		fmt.Println("ldx->", ldx, "stack", stack)
		for len(stack) > 0 {
			topIdx := stack[len(stack)-1]
			topVal := hArr[topIdx]
			fmt.Println("topVal->", topVal, "lval->", lval)

			if topVal < lval {
				stack = stack[:len(stack)-1]
			} else {
				nextHigh = topIdx
				break
			}

		}
		nextHigherIdx[ldx] = nextHigh
		stack.PushFront(ldx)
	}

	return nextHigherIdx
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

func TestNextHigher(t *testing.T) {
	arr := []int{4, 5, 3}
	nextHigher := computeIdxBound(arr)

	fmt.Println("next higher->", nextHigher)

}

func TestTrap(t *testing.T) {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	trap(arr)
}
