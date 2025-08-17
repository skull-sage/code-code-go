package adhoc_critical

// frequence bucketing
func topKFrequent(nums []int, k int) []int {

	freqMap := make(map[int]int, 0)
	for _, x := range nums {
		freqMap[x]++
	}

	bucketLen := len(nums)
	freqBucket := make([][]int, bucketLen+1, bucketLen+1)

	// there can be multiple x with same frequency f
	// bucketing all key(num) items that has same frequency
	for key, freq := range freqMap {
		if freqBucket[freq] == nil {
			freqBucket[freq] = make([]int, 0)
		}
		freqBucket[freq] = append(freqBucket[freq], key)
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
