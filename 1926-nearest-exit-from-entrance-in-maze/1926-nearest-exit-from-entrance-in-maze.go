func nearestExit(maze [][]byte, entrance []int) int {
    if maze == nil {
        return 0
    }
    
    q := [][]int{}
    m := len(maze)
    n := len(maze[0])
    
    
    // enqueue all the exits from all the edges
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == entrance[0] && j == entrance[1] {continue}
            if i == 0 || i == m-1 {
                if maze[i][j] == '.' {
                    q = append(q, []int{i,j})
                    maze[i][j] = '+'
                }
            } else if j == 0 || j == n-1 {
                if maze[i][j] == '.' {
                    q = append(q, []int{i,j})                    
                    maze[i][j] = '+'
                }
            }
        }
    }
    if len(q) == 0 {
        return -1
    }
    dirs := [][]int{ {1,0},{-1,0}, {0,1}, {0,-1}}
    level := 0
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq[0] == entrance[0] && dq[1] == entrance[1] {
                return level
            }
            for _, dir := range dirs {
                r := dq[0] + dir[0]
                c := dq[1] + dir[1]
                if r >= 0 && r < m && c >= 0 && c < n && maze[r][c] != '+' {
                    maze[r][c] = '+'
                    q = append(q, []int{r,c})
                }
            }
            qSize--
        }
        level++
    }
    return -1
}