func shortestPathBinaryMatrix(grid [][]int) int {
    if grid == nil  || len(grid) == 0 || len(grid[0]) == 0 || grid[0][0] != 0 || grid[len(grid)-1][len(grid[0])-1] != 0 {
        return -1
    }
    
    dirs := [][]int{  
        // horizontal and vertical
        {1,0},
        {-1,0},
        {0,1},
        {0,-1},

        // diagonal
        {1,1},
        {1,-1},
        {-1,1},
        {-1,-1},
    }
    m := len(grid)
    n := len(grid[0])
    q := [][]int{{0,0}}
    grid[0][0] = 1
    path := 0
    for len(q) != 0{
        
        qSize := len(q)
        // fmt.Println("Current queue: ", q, qSize)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            
            if dq[0] == m-1 && dq[1] == n-1 {
                return path+1
            }
            
            for _, dir := range dirs {
                r := dq[0] + dir[0]
                c := dq[1] + dir[1]
                if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == 0 {
                    grid[r][c] = 1
                    q = append(q, []int{r,c})
                }
            }
            qSize--
        }
        path++
        
    }
    
    return -1
}