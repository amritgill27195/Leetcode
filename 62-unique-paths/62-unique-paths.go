func uniquePaths(m int, n int) int {
    dr := m-1
    dc := n-1
    memo := make([][]int, m)
    for idx, _ := range memo {
        memo[idx] = make([]int, n)
    }
    var dfs func(r,c int) int
    dfs = func(r, c int) int {
        // base
        if r == m || r < 0 || c == n || c < 0 {return 0}
        if r == dr && c == dc {return 1}
        if memo[r][c] == 0 {
            memo[r][c] = dfs(r+1, c) + dfs(r,c+1)
        }
        return memo[r][c]
    }
    
    return dfs(0,0)
}