package basics

import (
	"fmt"
	"testing"
)

func traceSlice(slice []int, idx int) {
	if idx < 5 {
		slice[0], slice[idx] = slice[idx], slice[0]
		fmt.Printf("%p, %v\n", slice, slice)
		traceSlice(slice, idx+1)
	}
}

func TestSlicePass(t *testing.T) {
	slice := make([]int, 0, 5)

	fmt.Printf("%p, %v\n", slice, slice)
	for idx := range 5 {
		slice = append(slice, idx)
	}
	fmt.Printf("%p, %v\n", slice, slice)

	traceSlice(slice, 0)
}
