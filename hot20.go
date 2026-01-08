package hot

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	if n == 0 || m == 0 {
		return false
	}
	x, y := 0, n-1
	for x <= m && y >= 0 {
		if target == matrix[x][y] {
			return true
		}
		if target > matrix[x][y] {
			x++
		} else if target < matrix[x][y] {
			y--
		}

	}
	return false
}
