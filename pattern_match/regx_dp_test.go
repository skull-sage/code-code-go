package pattern_match

import (
	"fmt"
	"testing"
)

type Element struct {
	Val     byte
	hasStar bool
}

func simplifyPT(pattern string) []Element {
	pt := make([]Element, 0, len(pattern))
	for i := 0; i < len(pattern); i++ {
		val := pattern[i]
		hasStar := false
		if i+1 < len(pattern) && pattern[i+1] == '*' {
			hasStar = true
			i++
		}
		pt = append(pt, Element{val, hasStar})
	}

	return pt
}

func isMatch(text, pattern string) bool {
	txtLen := len(text)
	pt := simplifyPT(pattern)
	ptLen := len(pt)

	dpMatch := make([][]bool, txtLen+1, txtLen+1)
	for i := 0; i <= txtLen; i++ {
		dpMatch[i] = make([]bool, ptLen+1, ptLen+1)
		//fmt.Println(dpMatch[i])
	}

	dpMatch[0][0] = true

	for j := 1; j <= ptLen; j++ {
		ptItem := pt[j-1]
		if ptItem.hasStar {
			dpMatch[0][j] = dpMatch[0][j-1]
		} else {
			break
		}

	}
	fmt.Println(dpMatch[0])
	for i := 1; i <= txtLen; i++ {
		txtVal := text[i-1]
		for j := 1; j <= ptLen; j++ {
			ptItem := pt[j-1]

			if ptItem.Val == txtVal || ptItem.Val == '.' {
				dpMatch[i][j] = dpMatch[i-1][j-1]

			}
			if ptItem.hasStar {
				if ptItem.Val == txtVal || ptItem.Val == '.' {
					dpMatch[i][j] = dpMatch[i-1][j] || dpMatch[i][j-1]
				} else {
					dpMatch[i][j] = dpMatch[i][j-1]
				}
			}

		}
		fmt.Println(dpMatch[i])
	}

	return dpMatch[txtLen][ptLen]
}

func TestMatch(t *testing.T) {
	text := "aab"
	pattern := "aa*b*b"

	fmt.Println(isMatch(text, pattern))
}
