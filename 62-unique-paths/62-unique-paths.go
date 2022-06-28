func uniquePaths(m int, n int) int {
    dp := make([]int, n+1)
    dp[n-1] = 1
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            dp[j] = dp[j] + dp[j+1]
        }
    }
    return dp[0]
}