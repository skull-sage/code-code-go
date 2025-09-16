package adhoc_critical

import (
	"math/rand"
	"sort"
)

type Solution struct {
	prefixW []int
}

func NewSolution(w []int) Solution {
	prefixW := make([]int, len(w))
	prefixW[0] = w[0]
	for idx := 1; idx < len(w); idx++ {
		prefixW[idx] = prefixW[idx-1] + w[idx]
	}
	return Solution{
		prefixW: prefixW,
	}

}

func (this *Solution) PickIndex() int {
	totalSum := this.prefixW[len(this.prefixW)-1]
	randWeight := rand.Intn(totalSum) + 1
	randIdx := sort.SearchInts(this.prefixW, randWeight)
	return randIdx

}
