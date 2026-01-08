package hot

import "slices"

//O(MN),O(M+N)
func setZeroes1(matrix [][]int) {
	row := []int{}
	col := []int{}
	n := len(matrix)
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				row = append(row, i)
				col = append(col, j)
			}
		}
	}

	for i := 0; i < len(row); i++ {
		for j := 0; j < m; j++ {
			matrix[row[i]][j] = 0
		}
	}

	for i := 0; i < len(col); i++ {
		for j := 0; j < n; j++ {
			matrix[j][col[i]] = 0
		}
	}
}
//O(MN),O(1)
func setZeroes(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])
	fitstrowcontainszero := slices.Contains(matrix[0], 0)
	firstcolcontainszero := false
	for i := 0; i < n; i++ {
		if matrix[i][0] == 0 {
			firstcolcontainszero = true
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstcolcontainszero {
		for i := 0; i < n; i++ {
			matrix[i][0] = 0
		}
	}
	if fitstrowcontainszero {
		for i := 0; i < m; i++ {
			matrix[0][i] = 0
		}
	}
}
