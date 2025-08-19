package adhoc_critical

import (
	"fmt"
	"testing"
)

type Input struct {
	 
	lower int
	upper int
	count int
}

func sortAndCount(arrSlice []int, input *Input, level int) []int {


	if len(arrSlice) <=1 {

		if len(arrSlice) == 1 {
			input.count++
		}
		return arrSlice
	}
	 

	mid := len(arrSlice)/2 // l = 2, h = 7, m = 4

	leftSlice := arrSlice[0:mid]
	rightSlice := arrSlice[mid:len(arrSlice)]

	leftSlice = sortAndCount(leftSlice, input, level+1)  // 0-1
	rightSlice = sortAndCount(rightSlice, input, level+1) // 2-3

	
	return merge(leftSlice, rightSlice, level)

}

func countRange(leftSlice, rightSlice []int, input *Input){

	i := 0
	j := 0
	// a subset sum =>  ---arr[k]{-----}arr[i]----
	for k := 0; k < len(leftSlice); k++ {
		for i < len(rightSlice) && (rightSlice[i] - leftSlice[k]) < input.lower {
			i++
		}

		j = i
		//fmt.Printf("j=%d, k=%d %d", j, k, (this.pfSum[j] - this.pfSum[k]))
		for j < len(rightSlice) && (rightSlice[j] - leftSlice[k]) <= input.upper {
			//fmt.Printf("j=%d, k=%d %d", j, k, (this.pfSum[j] - this.pfSum[k]))
			j++
		}

	}

	//fmt.Println(low, mid, high, "j", j, "i", i)
	input.count += (j - i)
}

func merge(leftSlice []int, rightSlice []int, level int) []int {
	//merge step

	mergeCopy := make([]int, len(leftSlice)+len(rightSlice))

	i := 0
	j := 0
	k := 0
	for i < len(leftSlice) && j < len(rightSlice) {
		if leftSlice[i] <= rightSlice[j] {
			mergeCopy[k] = leftSlice[i]
			k++
			i++
		} else {
			mergeCopy[k] = rightSlice[j]
			k++
			j++
		}

	}

	for i < len(leftSlice) {
		mergeCopy[k] = leftSlice[i]
		k++
		i++

	}

	for j < len(rightSlice) {
		mergeCopy[k] = rightSlice[j]
		k++
		j++
	}

	fmt.Println("# after merging:", mergeCopy)
	return mergeCopy
}

func countRangeSum(nums []int, lower int, upper int) int {

	pfSum := make([]int, len(nums))

	pfSum[0] = nums[0]

	for idx := 1; idx < len(nums); idx++ {
		pfSum[idx] = pfSum[idx-1] + nums[idx]
	}

	fmt.Println("before sorting", pfSum)
	input := Input{lower, upper, 0}
	pfSum = sortAndCount(pfSum, &input, 0)
	fmt.Println("after sorting", pfSum)
	return input.count

}

func TestRangeCount(t *testing.T) {

	nums := []int{-1}
	lower := -2
	upper := 5

	fmt.Println(countRangeSum(nums, lower, upper))
}
