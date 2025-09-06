package range_query

import (
	"fmt"
	"sort"
)

func subarraySum(nums []int, k int) int {
	idxListMap := make(map[int][]int, 0)

	psum := make([]int, len(nums)+1)
	psum[0] = 0
	for idx := 1; idx < len(psum); idx++ {
		sum := psum[idx-1] + nums[idx-1]
		psum[idx] = sum
		if idxListMap[sum] == nil {
			idxListMap[sum] = make([]int, 0)
		}
		idxListMap[sum] = append(idxListMap[sum], idx)
	}

	fmt.Println(psum)
	fmt.Println(idxListMap)

	result := 0
	for idx := 0; idx < len(psum); idx++ {

		target := psum[idx] + k
		list := idxListMap[target]

		idxStart := sort.SearchInts(list, idx)
		result += (len(list) - idxStart)
	}

	return result

}
