package adhoc_critical

import (
	"fmt"
	"testing"
)

type Input struct {
	pfSum []int
	lower int
	upper int
	count int
}

func (this *Input) sortAndCount(low, high, level int) {

	if low >= high {
		return
	}

	mid := low + (high-low)/2 // l = 2, h = 7, m = 4

	this.sortAndCount(low, mid, level+1)
	this.sortAndCount(mid+1, high, level+1)

	i := mid + 1
	j := mid + 1
	// a subset sum =>  ---arr[k]{-----}arr[i]----
	for k := low; k <= mid; k++ {
		for i <= high && (this.pfSum[i]-this.pfSum[k]) < this.lower {
			i++
		}

		//j = i
		//fmt.Printf("j=%d, k=%d %d", j, k, (this.pfSum[j] - this.pfSum[k]))
		for j <= high && (this.pfSum[j]-this.pfSum[k]) <= this.upper {
			//fmt.Printf("j=%d, k=%d %d", j, k, (this.pfSum[j] - this.pfSum[k]))
			j++
		}

	}

	//fmt.Println(low, mid, high, "j", j, "i", i)
	this.count += (j - i)
	this.merge(low, mid, high, level)

}

func (this *Input) merge(low, mid, high, level int) {
	//merge step

	left := make([]int, (mid + 1 - low)) //this.pfSum[low : mid+1]
	copy(left, this.pfSum[low:mid+1])

	right := make([]int, (high - mid))
	copy(right, this.pfSum[mid+1:high+1])

	fmt.Println("level", level, "l:", low, "m:", mid, "h:", high, "=> left", left, "right", right)
	i := 0
	j := 0
	k := 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			this.pfSum[k] = left[i]
			k++
			i++
		} else {
			this.pfSum[k] = right[j]
			k++
			j++
		}

		fmt.Println("# merging: ", i, j, k, this.pfSum)
	}

	for i < len(left) {
		this.pfSum[k] = left[i]
		k++
		i++

	}

	fmt.Println("# merging: ", i, j, k, this.pfSum)

	for j < len(right) {
		this.pfSum[k] = right[j]
		k++
		j++
	}

	fmt.Println("# after merging:", this.pfSum)
}

func countRangeSum(nums []int, lower int, upper int) int {

	pfSum := make([]int, len(nums))

	pfSum[0] = nums[0]

	for idx := 1; idx < len(nums); idx++ {
		pfSum[idx] = pfSum[idx-1] + nums[idx]
	}

	fmt.Println("before sorting", pfSum)
	input := Input{pfSum, lower, upper, 0}
	input.sortAndCount(0, len(pfSum)-1, 0)
	fmt.Println("after sorting", input.pfSum)
	return input.count

}

func TestRangeCount(t *testing.T) {
	nums := []int{-2, 5, -1}
	lower := -2
	upper := 5

	fmt.Println(countRangeSum(nums, lower, upper))
}
