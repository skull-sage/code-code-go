package range_query

type NumMatrix struct {
	psum [][]int
	rowN int
	colN int
}

func (this NumMatrix) translet(i, j int) int {
	return (i*this.colN + j) + 1 // the extra one to make psum index 1 based instead of zero
}

func Constructor(matrix [][]int) NumMatrix {
	rowN := len(matrix)
	colN := len(matrix[0])

	psum := make([][]int, rowN+1)

	for idx := range psum {
		psum[idx] = make([]int, colN+1)
	}

	psum[0][0] = 0

	for i := 1; i <= rowN; i++ {
		for j := 1; j <= colN; j++ {
			psum[i][j] = matrix[i-1][j-1] + psum[i-1][j] + psum[i][j-1] - psum[i-1][j-1]
		}
	}

	return NumMatrix{psum, rowN, colN}
}

func (this *NumMatrix) SumRegion(r1 int, c1 int, r2 int, c2 int) int {

	// the reason of adding r1c1 is because both area r2c1 and r1c2 contains area r1c1
	// we want to subtract it once
	return this.psum[r2+1][c2+1] - this.psum[r2+1][c1+1] - this.psum[r1+1][c2+1] + this.psum[r1+1][c1+1]

}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
