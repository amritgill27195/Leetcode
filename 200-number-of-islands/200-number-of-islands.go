func numIslands(grid [][]byte) int {
 
    if grid == nil {
        return 0
    }
    count := 0
    q := [][]int{}
    m := len(grid)
    n := len(grid[0])
    dirs := [][]int{ {0,-1}, {0,1}, {-1, 0}, {1,0} }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                q = append(q, []int{i,j})    
                grid[i][j] = '0' // mark this land visited
                for len(q) != 0 {
                    dq := q[0]
                    q = q[1:]
                    for _, dir := range dirs {
                        r := dq[0] + dir[0]
                        c := dq[1] + dir[1]
                        if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == '1' {
                            grid[r][c] = '0'
                            q = append(q, []int{r,c})
                        }
                    }
                }
                count++
            }
        }
    }
    return count 
}