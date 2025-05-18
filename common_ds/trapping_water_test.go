package common_ds

import (
	"fmt"
	"testing"
)

func trap(hArr []int) int {
	// find the max height
	right := len(hArr) - 1
	trapped := 0
	tmpTrapped := 0

	for left := right; left >= 0; left-- {
		if hArr[left] >= hArr[right] {
			fmt.Println("left: ", left, " right: ", right, " tmpTrapped: ", tmpTrapped)
			right = left
			trapped += tmpTrapped
			tmpTrapped = 0
			fmt.Println("trapped: ", trapped)
		} else if hArr[left] < hArr[right] {
			tmpTrapped += hArr[right] - hArr[left]
		}
	}

	return trapped

}

func TestTrap(t *testing.T) {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	trap(arr)
}
