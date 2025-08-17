package adhoc_critical

type Input struct {
	pfSum []int
	lower int
	upper int
	count int
}

func (this *Input) sortAndCount(low, high int) {

	if low >= high {
		return
	}

	mid := (high + low) / 2 // l = 2, h = 7, m = 4

	this.sortAndCount(low, mid)
	this.sortAndCount(mid, high)

	arr := this.pfSum

	i := mid
	j := mid
	// a subset sum =>  ---arr[k]{-----}arr[i]----
	for k := 0; k < mid; k++ {
		for i < high && (arr[i]-arr[k]) < this.lower {
			i++
		}

		for j < high && (arr[j]-arr[k]) <= this.upper {
			j++
		}

	}

	this.count += (j - i)
	this.merge(low, mid, high)

}

func (this *Input) merge(low, mid, high int) {
	//merge step

	lowEnd := this.pfSum[low:mid]
	highEnd := this.pfSum[mid:high]

	i := 0
	j := 0
	k := 0
	for i < mid && j < high {
		if lowEnd[i] <= highEnd[j] {
			this.pfSum[k] = lowEnd[i]
			k++
			i++
		} else {
			this.pfSum[k] = highEnd[j]
			k++
			j++
		}
	}

	for i < mid {
		this.pfSum[k] = lowEnd[i]
		k++
		i++
	}

	for j < high {
		this.pfSum[k] = highEnd[j]
		k++
		j++
	}
}

func countRangeSum(nums []int, lower int, upper int) int {

	pfSum := make([]int, len(nums))

	for idx := 0; idx < len(nums); idx++ {
		pfSum[idx] = pfSum[idx-1] + nums[idx]
	}

	input := Input{pfSum, lower, upper, 0}
	input.sortAndCount(0, len(pfSum))
	return input.count

}
