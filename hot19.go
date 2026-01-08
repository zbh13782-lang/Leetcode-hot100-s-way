package hot

import "math"

var dire = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func spiralOrder(matrix [][]int) []int {
	i, j, di := 0, 0, 0
	m, n := len(matrix), len(matrix[0])
	res := make([]int, m*n)
	cnt := m * n
	cur := 0
	for cur < cnt {
		res[cur] = matrix[i][j]
		matrix[i][j] = math.MaxInt
		x, y := i+dire[di][0], j+dire[di][1]
		if x >= m || y >= n || x < 0 || y < 0 || matrix[x][y] == math.MaxInt {
			di = (di + 1) % 4
		}
		i, j = i+dire[di][0], j+dire[di][1]
	}
	return res
}
