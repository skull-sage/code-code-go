package adhoc_critical

// frequence bucketing
func topKFrequent(nums []int, k int) []int {

	freqMap := make(map[int]int, 0)
	for _, x := range nums {
		freqMap[x]++
	}

	bucketLen := len(nums)
	freqBucket := make([][]int, bucketLen+1, bucketLen+1)

	// all items vals > 0 so bucketLen + 1
	for key, val := range freqMap {
		if freqBucket[val] == nil {
			freqBucket[val] = make([]int, 0)
		}
		freqBucket[val] = append(freqBucket[val], key)
	}

	result := make([]int, 0)
	for idx := bucketLen; idx >= 1; idx-- {
		for _, item := range freqBucket[idx] {
			result = append(result, item)
			if len(result) == k {
				return result
			}
		}
	}

	return result // required by go

}
