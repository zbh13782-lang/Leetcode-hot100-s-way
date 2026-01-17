package hot

type pair struct {
	x, y int
}

var dicret = []pair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func orangesRotting(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	raw := []pair{}
	fresh := 0
	for i, row := range grid {
		for j, value := range row {
			if value == 1 {
				fresh++
			}
			if value == 2 {
				raw = append(raw, pair{i, j})
			}
		}
	}

	for fresh > 0 && len(raw) != 0 {
		cur := raw
		raw = []pair{}
		ans++
		for _, v := range cur {

			for _, d := range dicret {
				x := v.x + d.x
				y := v.y + d.y
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					grid[x][y] = 2
					fresh--
					raw = append(raw, pair{x, y})
				}
			}
		}
	}
	if fresh > 0 {
		ans = -1
		return
	}
	return
}
