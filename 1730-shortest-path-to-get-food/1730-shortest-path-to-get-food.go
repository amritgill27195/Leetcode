
func getFood(grid [][]byte) int {

	if grid == nil {
		return -1
	}

	var endPosition []int
	m := len(grid)
	n := len(grid[0])
	var q [][]int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '*' {
				endPosition = append(endPosition, i, j)
			} else if grid[i][j] == '#' {
				q = append(q, []int{i, j})
			}
		}
	}
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	if len(endPosition) == 0 || len(q) == 0 {
		return -1
	}

	level := 0
	for len(q) != 0 {
		qSize := len(q)

		for qSize != 0 {
			dq := q[0]
			q = q[1:]

			if dq[0] == endPosition[0] && dq[1] == endPosition[1] {
				return level
			}

			for _, dir := range dirs {
				r := dq[0] + dir[0]
				c := dq[1] + dir[1]
				if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] != 'X' {
					grid[r][c] = 'X'
					q = append(q, []int{r, c})
				}
			}
			qSize--
		}
		level++
	}
	return -1
}
