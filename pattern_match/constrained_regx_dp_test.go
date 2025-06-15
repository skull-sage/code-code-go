package pattern_match

import (
	"fmt"
	"testing"
)

// while most of leetcode seems to have n^2 DP solution
// Pressumabely if problem didn't require expression match with "".*"
// there is linear time solution exist

type PatItem struct {
	val        byte
	hasStar    bool
	leastCount int
}

type TextElm struct {
	val   byte
	count int
}

func simplifyPattern(pattern string) []PatItem {
	m := len(pattern)
	pt := make([]PatItem, 0, m)

	for i := 0; i < m; {
		val := pattern[i]
		hasStar := false
		leastCount := 0

		j := i
		for j < m && pattern[j] == val {

			if j+1 < m && pattern[j+1] == '*' {
				hasStar = true
				j = j + 2

			} else {
				j = j + 1
				leastCount++
			}
		}
		pt = append(pt, PatItem{val, hasStar, leastCount})
		i = j

	}

	return pt
}

func simplifyText(text string) []TextElm {
	n := len(text)
	tt := make([]TextElm, 0, n)

	for i := 0; i < n; {

		val := text[i]
		j := i
		for j < n && text[j] == val {
			j++
		}
		tt = append(tt, TextElm{val, (j - i)})
		i = j
	}

	return tt

}

// i write my own way to solve
func isMatchLinear(text string, pattern string) bool {
	pt := simplifyPattern(pattern)
	tt := simplifyText(text)

	j := 0
	i := 0
	for i < len(tt) {

		// there are less pattern TextElms than text TextElms
		if j == len(pt) {
			fmt.Println("#text has extra TextElms comparing to Pattern")
			return false
		}

		ttItem := tt[i]
		ptItem := pt[j]

		fmt.Println("# comparing: ", ttItem, "to pt-Item", ptItem)
		if ttItem.val != ptItem.val {
			if ptItem.hasStar == false {
				return false
			} else if ptItem.leastCount > 0 {
				return false
			} else { // not matched but has * and leastCount == 0 so skip to next
				j++
			}
		} else if ptItem.leastCount > ttItem.count {
			return false
		} else if ptItem.leastCount == ttItem.count || ptItem.hasStar {
			// we finally has pattern match for i-th value, move to next
			i++
			j++
		}

	}

	if i == len(tt) {
		for j < len(pt) {
			if pt[j].hasStar {
				j++
			} else {
				fmt.Println("#pattern has extra not-skippable TextElms")
				return false
			}
		}
	}

	return true
}

func TestIsMatch(t *testing.T) {

	text := "ab"
	pattern := ".*"
	fmt.Println("isMatch: ", isMatch(text, pattern))
	//fmt.Println(simplifyPattern("aaa*bbc*"))
	//fmt.Println(simplifyPattern("aaa*"))
	//fmt.Println(simplifyPattern("a*aaa*aa*"))
	//fmt.Println(simplifyPattern("a*"))

}
