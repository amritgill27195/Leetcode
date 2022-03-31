func orangesRotting(grid [][]int) int {
    if grid == nil {
        return 0
    }
    queue := [][]int{}
    m := len(grid)
    n := len(grid[0])
    fresh := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 2 {
                queue = append(queue, []int{i,j})
            } else if grid[i][j] == 1 { fresh++ }
        }
    }
    
    // if we dont have any fresh to rotten, no need to proceed even if we have rotten ones already
    if fresh == 0 {
        return 0
    }
    
    dirs := [][]int{{0,1},{0,-1},{-1,0},{1,0}}
    
    // bfs 
    time := 0
    for len(queue) != 0 {
        qSize := len(queue)
        for qSize != 0 {
            dq := queue[0]
            queue = queue[1:]
            
            for _, dir := range dirs {
                r := dq[0] + dir[0]
                c := dq[1] + dir[1]
                if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == 1 {
                    grid[r][c] = 2
                    fresh--
                    queue = append(queue, []int{r,c})
                }
            }
            qSize--
        }
        time++  
    }
    if fresh != 0 {
        return -1 // not possible to rotten all
    }
    return time-1
}