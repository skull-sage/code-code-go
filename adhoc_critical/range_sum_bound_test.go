package adhoc_critical

func sortAndCount(pfSum *[]int, low, high int) int {

	if low >= high {
		return 0
	}

	mid := (high + low)/2 // l = 2, h = 7, m = 4 

	sortAndCount(pfSum, low, mid)
	sortAndCount(pfSum, mid, high)

	count :=0
	i:= mid;
	
	for k:=0; k < mid; k++ {
		for i < high && (*pfSum[i] - pfSum[k]) < low {
			i++
		}
	} 


	merge(pfSum, low, mid, high)

	return count

}

func merge(pfSum *[]int, low, mid, high int){
	//merge step
	arr := *pfSum	
	lowEnd := arr[low:mid];
	highEnd := arr[mid:high];

	i:=0
	j:=0
	k := 0
	for i < mid && j < high {
		if lowEnd[i] <= highEnd[j]{
			arr[k] = lowEnd[i]
			k++
			i++ 
		} else {
			arr[k] = highEnd[j];
			k++
			j++
		}
	}

	for i < mid {
		arr[k] = lowEnd[i]
		k++
		i++ 
	}

	for j < high {
		arr[k] = highEnd[j]
		k++
		j++
	}
}

func countSumRange(nums []int, idxDiff int, valDiff int) int {

	pfSum := make([]int, len(nums));

	for idx, x := range nums {
		pfSum[idx] = pfSum[idx-1] + x;	
	}

	return sortAndCount(&pfSum, 0, len(pfSum))
 
}