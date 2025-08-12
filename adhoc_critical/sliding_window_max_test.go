package adhoc_critical

import (
	"container/list"
	"math"
)

type IdxTuple struct {
	val int
	idx int
}

func maxSlidingWindow(nums []int, k int) []int {

	if k == 1 {
		return nums
	}

	n := len(nums)
	max := -math.MaxInt32 //IdxTuple{val:-math.MaxInt32, idx:0}

	ll := list.New()
	result := make([]int, 0)

	for idx := k - 1; idx >= 0; idx-- {
		num := nums[idx]

		if max <= num {
			max = num
			ll.PushFront(idx)
		}

	}

	result = append(result, max)
	//fmt.Println(ll)

	for idx := k; idx < n; idx++ {
		valK := nums[idx]

		e := ll.Front()
		for e != nil {
			tIdx := e.Value.(int)
			if tIdx <= (idx - k) {
				eNext := e.Next()
				ll.Remove(e)
				e = eNext
			} else {
				break
			}

		}

		e = ll.Back()
		for e != nil {
			tIdx := e.Value.(int)
			tVal := nums[tIdx]

			if tVal <= valK {
				ePrev := e.Prev()
				ll.Remove(e)
				e = ePrev
			} else {
				break
			}

		}

		ll.PushBack(idx)
		maxIdx := ll.Front().Value.(int)

		result = append(result, nums[maxIdx])

	}

	return result

}
