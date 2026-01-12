package hot

func numIslands(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])
	var infect func(int, int)
	infect = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		infect(i+1, j)
		infect(i-1, j)
		infect(i, j+1)
		infect(i, j-1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ans++
				infect(i, j)
			}
		}
	}
	return
}
