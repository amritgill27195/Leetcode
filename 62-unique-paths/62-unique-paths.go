func uniquePaths(m int, n int) int {
    dp := make([][]int, m+1)
    for idx, _ := range dp {
        dp[idx] = make([]int, n+1)
    }
    // there is only 1 way to get out from the last cell ( which is right ) therefore only 1 way
    dp[m-1][n-1] = 1
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            if i == m-1 && j == n-1 {continue}
            dp[i][j] = dp[i+1][j] + dp[i][j+1]
        }
    }
    return dp[0][0]
}