package common_ds

import (
	"fmt"
	"testing"
)

func trap(hArr []int) int {
	// find the max height

	trapped := 0
	maxHeight := hArr[0]
	maxHeightIdx := 0

	for idx := 1; idx < len(hArr); idx++ {
		if hArr[idx] > maxHeight {
			maxHeight = hArr[idx]
			maxHeightIdx = idx
		}
	}

	// left to right scan
	for ldx := 0; ldx < maxHeightIdx; {
		tmpTrapped := 0

		start := hArr[ldx]
		rdx := ldx + 1

		for ; rdx < maxHeightIdx && hArr[rdx] < start; rdx++ {
			tmpTrapped += (start - hArr[rdx])
		}
		ldx = rdx
		//fmt.Println("start:", start, "end:", hArr[ldx], "tmpTrapped:", tmpTrapped)
		trapped += tmpTrapped

	}

	for rdx := len(hArr) - 1; rdx > maxHeightIdx; {

		tmpTrapped := 0
		start := hArr[rdx]
		ldx := rdx - 1
		for ; ldx > maxHeightIdx && hArr[ldx] < start; ldx-- {
			tmpTrapped += (start - hArr[ldx])
		}
		rdx = ldx
		trapped += tmpTrapped
	}

	return trapped
}

func TestTrap(t *testing.T) {
	arr := []int{2, 3, 4}
	trapped := trap(arr)
	fmt.Println("Trapped ", trapped)
}
