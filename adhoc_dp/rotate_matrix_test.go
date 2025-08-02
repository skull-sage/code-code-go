package adhoc_dp

import (
	"fmt"
	"testing"
)

// rotate 90 : transpose then reverse row
// rotate 180 : reverse row then reverse col
// rotate 270 : transpose then reverse col

// transpose then reverse
func rotate90(matrix [][]int) [][]int {
	transpose(matrix)
	n := len(matrix)
	for i := range matrix {
		for j := 0; j < n/2; j++ {
			r := n - 1 - j
			matrix[i][j], matrix[i][r] = matrix[i][r], matrix[i][j]
		}
	}

	return matrix
}

// 1   2   3   4 => 1 4 7
// 5   6   7   8 => 2 5 8
// 9  10  11  12 => 3 6 9
// 13  14  15  16 =>
func transpose(matrix [][]int) [][]int {
	n := len(matrix)

	for i := 0; i < n; i++ {

		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}

	}

	return matrix
}

func TestTranspose(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	//transpose := transpose(matrix)
	rotate := rotate90(matrix)
	fmt.Println(rotate)
}
