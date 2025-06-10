package pattern_match

import (
	"fmt"
	"testing"
)

func KMPMatch(T, P string) {
	n := len(T)
	m := len(P)
	prf := computePrefix(P)
	k := 0
	for i := 1; i <= n; i++ {
		for k > 0 && P[k] != T[i-1] {
			k = prf[k]
		}

		if P[k] == T[i-1] {
			k = k + 1
		}

		if k == m {
			fmt.Println("Pattern Occurs with Shift: ", i-m)
			k = prf[k]
		}
	}
}

// understanding the prefix function is the holy grail for KMP matcher
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
	P := "abcabcabcabc"
	computePrefix(P)
}
