func searchMatrix(matrix [][]int, target int) bool {
	// as its a sorted arr if target is smaller than first or bigger than last el, it wont be in the matrix
	// so initial check could speed it up
	if len(matrix) == 0 || len(matrix[0]) == 0 {
    return false
}
if target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1] {
    return false
}
	rows, cols := len(matrix), len(matrix[0])
	l, r := 0, (rows * cols) - 1 // left at 0, right at (rowslen * colslen) - 1
	for l <= r {
	 m := (l+r)/2
	 // m is from 0 to n where n is rows * cols - 1
	 // to map it to row(index of inner arrs) div by cols num(len of each inner arr)
	 // m % cols gives pos inside inner arr
	 rowPos := m/cols 
	 colPos := m%cols
	 if target == matrix[rowPos][colPos] {
		return true
	 } else if target > matrix[rowPos][colPos] {
		l = m+1 // update left ptr if its bigger
	 } else {
		r = m-1
	 }
	}
	return false
}
