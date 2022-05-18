func shortestPathBinaryMatrix(grid [][]int) int {
    level := 0
    m := len(grid)
    n := len(grid[0])
    destRow := m-1
    destCol := n-1
    if grid[destRow][destCol] != 0 || grid[0][0] != 0 {
        return -1
    }
    dirs := [][]int{
        {1,0},
        {-1,0},
        {0,1},
        {0,-1},
        {-1,-1},
        {-1,1},
        {1,-1},
        {1,1},
    }
    grid[0][0] = 1
    q := [][]int{{0,0}}
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            
            if dq[0] == destRow && dq[1] == destCol {
                return level+1
            }
            for _, dir := range dirs {
                r := dir[0] + dq[0]
                c := dir[1] + dq[1]
                if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == 0 {
                    grid[r][c] = 1
                    // to exit early ... 
                    if r == destRow && c == destCol {return level+2}
                    q = append(q, []int{r,c})
                }
            }
            qSize--
        }
        level++
    }
    return -1
}