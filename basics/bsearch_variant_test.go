package basics

import (
	"fmt"
	"testing"
)

/**
there are five variants of a binary search
- does the target exist
- target exist and it's first occurance
- target exist and it's last occurance
- greatest element less than target
- smallest element larger than target
*/

/*
* I intend to find a single variant
 */

func binaryTarget(nums []int, target int) {
	low := 0
	high := len(nums) - 1
	targetIdx := -1
	for low <= high {
		mid := low + (high-low)/2
		midVal := nums[mid]

		if target < midVal {
			high = mid - 1
		} else if target > midVal {
			low = mid + 1
		} else if target == midVal {
			targetIdx = mid
			//low = mid + 1 //for smallest > target
			high = mid - 1 // for greatest < target

		}
	}

	fmt.Println("#searching", target)
	fmt.Println("- result: ", "low", low, "high", high, "targetIdx", targetIdx, "gr-smaller", nums[low-1])

	// in a nutshell
	//  in case of missing target, low is the index if we were to insert it into search
	//	otherwise, we use target is the index
}

func TestBinarySearch(t *testing.T) {
	arrList := []int{1, 3, 5, 6}
	fmt.Println("#input:", arrList)
	// Lets try for missing targets: 1, 4, 8
	//binaryTarget(arrList, 1)
	//binaryTarget(arrList, 4)
	//binaryTarget(arrList, 8)

	// try a target that exist
	binaryTarget(arrList, 5)

}
