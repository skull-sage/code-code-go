package pattern_match

import (
	"fmt"
	"testing"
)

func computePrefix(P string) []int {
	m := len(P)
	prf := make([]int, m, m)
	k := 0
	for q := 1; q < m; q++ {
		fmt.Printf("#itr: P[q=%d] = %c, P[k=%d] = %c\n", q, P[q], k, P[k])
		// skip finding prefix as for k = 0 as ther is no match
		for k > 0 && P[k] != P[q] {
			k = prf[k]
		}
		if P[k] == P[q] {
			k = k + 1
		}

		prf[q] = k

		fmt.Println("      result:", prf)
		fmt.Println()
	}

	return prf
}

func TestKMP(t *testing.T) {
	P := "abacaba"
	computePrefix(P)
}
