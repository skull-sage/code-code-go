package pattern_match

import (
	"fmt"
	"testing"
)

func isMatch(text string, pattern string) bool {
	n := len(text)
	m := len(pattern)

	dpArr := make([][]bool, (n+1)*(m+1))

	for i := 0; i < n+1; i++  {
        if (p[i] == '*' && dp[0][i-1]) {
            dp[0][i+1] = true;
        }
    }
	
    for (int i = 0 ; i < s.length(); i++) {
        for (int j = 0; j < p.length(); j++) {
            if (p.charAt(j) == '.') {
                dp[i+1][j+1] = dp[i][j];
            }
            if (p.charAt(j) == s.charAt(i)) {
                dp[i+1][j+1] = dp[i][j];
            }
            if (p.charAt(j) == '*') {
                if (p.charAt(j-1) != s.charAt(i) && p.charAt(j-1) != '.') {
                    dp[i+1][j+1] = dp[i+1][j-1];
                } else {
                    dp[i+1][j+1] = (dp[i+1][j] || dp[i][j+1] || dp[i+1][j-1]);
                }
            }
        }
    }
    return dp[s.length()][p.length()]; 
}

func TestIsMatch(t *testing.T) {

}
