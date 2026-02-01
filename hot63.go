package hot

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	var oTot func(num int) (row, col int)
	
	oTot = func(num int) (row int, col int) {
		row = num / n
		col = num % n
		return
	}
	

	var isgreen func(num, target int) bool
	isgreen = func(n, t int) bool {
		return n >= t
	}

	alllen := m * n

	left, right := -1, alllen
	x, y := 0, 0
	for right-left > 1 {
		middle := left + (right-left)/2
		x, y = oTot(middle)
		if isgreen(matrix[x][y], target) {
			right = middle
		} else {
			left = middle
		}
	}

	x, y = oTot(right)
	return matrix[x][y] == target
}
