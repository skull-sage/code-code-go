package adhoc_critical

import (
	"fmt"
	"testing"
)

type Input struct {
	 
	lower int64
	upper int64
	count int
}

func sortAndCount(arrSlice []int64, input *Input, level int) []int64 {


	if len(arrSlice) <=1 {

		if len(arrSlice) == 1 {
            if arrSlice[0] >= input.lower && arrSlice[0] <= input.upper {
			    input.count++
               
            }
		}
		return arrSlice
	}
	 

	mid := len(arrSlice)/2 // l = 2, h = 7, m = 4

	leftSlice := arrSlice[0:mid]
	rightSlice := arrSlice[mid:len(arrSlice)]

	leftSlice = sortAndCount(leftSlice, input, level+1)  // 0-1
	rightSlice = sortAndCount(rightSlice, input, level+1) // 2-3

    countRange(leftSlice, rightSlice, input)
	
	return merge(leftSlice, rightSlice, level)

}

func countRange(leftSlice, rightSlice []int64, input *Input){
 

	
	for i := 0; i < len(leftSlice); i++ {
		
		// rightSlice[j1] >= leftSlice[i] + input.Lower
		// rightSlice[j2] <= leftSlice[i] + input.Upper

		for  j < len(rightSlice) && rightSlice[j] - leftSlice[i] < input.lower {
			j++
		}
		for  j < len(rightSlice) && rightSlice[j] - leftSlice[i] <= input.upper {
			input.count++
			j++
		}
		 
		
	}
    
}

func merge(leftSlice []int64, rightSlice []int64, level int) []int64 {
	//merge step

	mergeCopy := make([]int64, len(leftSlice)+len(rightSlice))

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

	//fmt.Println("#level", level, "=> merging:", mergeCopy)
	return mergeCopy
}

func countRangeSum(nums []int, lower int, upper int) int {

	pfSum := make([]int64, len(nums))

	pfSum[0] = int64(nums[0])

	for idx := 1; idx < len(nums); idx++ {
		pfSum[idx] = pfSum[idx-1] + int64(nums[idx])
	}

	fmt.Println("before sorting", pfSum)
	input := Input{int64(lower), int64(upper), 0}
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
