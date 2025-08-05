package basics

import (
	"fmt"
	"testing"
)

// the slice arg is a copy of the passed slice, hence any change to it are bounded by function traceSlice
// the change won't be reflected in the calling function that passed the argment slice :: []int
func traceSlice(slice []int, idx int) {
	if idx < 5 {
		slice = append(slice, idx*2)
		fmt.Printf("%p, %v\n", slice, slice)
		traceSlice(slice, idx+1)
	}
}

func traceSliceWithPtr(slice *[]int, idx int) {
	if idx < 5 {
		*slice = append(*slice, idx*2)
		fmt.Printf("%p, %v\n", slice, slice)
		traceSlice(*slice, idx+1)
	}
}

func TestSlicePass(t *testing.T) {
	slice := make([]int, 0, 5)

	fmt.Printf("%p, %v\n", slice, slice)
	for idx := range 5 {
		slice = append(slice, idx)
	}

	traceSlice(slice, 1)
	fmt.Printf("After Passing a copy: %p, %v\n", slice, slice)
	traceSliceWithPtr(&slice, 1)
	fmt.Printf("After Passing a ref: %p, %v\n", slice, slice)
}
