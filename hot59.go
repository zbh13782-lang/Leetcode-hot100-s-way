package hot

type pairs struct {
	x, y int
}

var directs = []pairs{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	l := len(word)
	ans := false
	var dfs func(int, int, int)
	dfs = func(i, j, cnt int) {
		if cnt == l {
			ans = true
			return
		}
		if i >= m || i < 0 || j >= n || j < 0 || board[i][j] != word[cnt] {
			return
		}
		temp := board[i][j]
		board[i][j] = '#'
		for _, direct := range directs {
			dx, dy := i+direct.x, j+direct.y
			dfs(dx, dy, cnt+1)
		}
		board[i][j] = temp
	}
	for i, row := range board {
		for j, _ := range row {
			dfs(i, j, 0)
			if ans {
				return ans
			}
		}
	}
	return ans
}
