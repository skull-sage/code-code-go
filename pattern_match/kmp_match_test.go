package pattern_match

import (
	"fmt"
	"testing"
)

func computePrefix(P string) []int {
	m := len(P)
	prf := make([]int, m+1, m+1)
	k := 0
	for q := 2; q <= m; q++ {
		fmt.Printf("#itr: P[q=%d] = %c, P[k=%d] = %c\n", q-1, P[q-1], k, P[k])
		// skip finding prefix as for k = 0 as ther is no match
		for k > 0 && P[k] != P[q-1] {
			k = prf[k]
		}
		if P[k] == P[q-1] {
			k = k + 1
		}

		prf[q] = k

		fmt.Println("      result:", "k=", k, prf[1:])
		fmt.Println()
	}

	return prf
}

func TestKMP(t *testing.T) {
	//P := "abacaba"
	P := "abcdabcdabab"
	computePrefix(P)
}
