func getMaximumGold(grid [][]int) int {
    maxGold := 0
    m := len(grid)
    n := len(grid[0])
    dirs := [][]int{{1,0},{-1,0},{0,-1},{0,1}}
    var dfs func(sum, r, c int) 
    dfs = func(sum, r, c int) {
        // base
        if r < 0 || r == m || c < 0 || c == n || grid[r][c] <= 0 {
            return
        }
        sum += grid[r][c]
        if sum > maxGold {maxGold = sum}
        grid[r][c] *= -1
        // logic
        for _, dir := range dirs {
            dfs(sum, r+dir[0], c+dir[1])
        }
        grid[r][c] *= -1        
    }
    
    for i := 0;i < m; i++ {
        for j := 0;j < n; j++ {
            if grid[i][j] != 0 {
                dfs(0,i, j)
            }
        }
    }
    return maxGold
}

func max(x, y int) int {
    if x > y {return x}
    return y
}