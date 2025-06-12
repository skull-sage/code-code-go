package pattern_match

import (
	"testing"
)

type Element struct {
	val        byte
	hasStar    bool
	leastCount int
}

func simplifyPattern(pattern string) []Element {
	m := len(pattern)
	pt := make([]Element, 0, m)

	for i := 0; i < m; {
		val := pattern[i]
		hasStar := false
		leastCount := 0

		j := i + 1
		for j < m && pattern[j] == val {

			if pattern[j+1] == '*' {
				hasStar = true
				j = j + 2

			} else {
				j = j + 1
				leastCount++
			}
		}
		pt = append(pt, Element{val, hasStar, leastCount})
		i = j

	}

	return pt
}

// i write my own way to solve
func isMatch(text string, pattern string) bool {
	pt := simplifyPattern(pattern)
	txt := simplifyText(text)

}

func TestIsMatch(t *testing.T) {

	var text string
	var pattern string

	text = "aa"
	pattern = "a"

	if isMatch(text, pattern) {
		t.Errorf("expect false, but got true")
	}
}
