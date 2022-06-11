func change(amount int, coins []int) int {
    m := len(coins)
    n := amount
    dp := make([][]int, m+1)
    for idx, _ := range dp {
        dp[idx] = make([]int, n+1)
        dp[idx][0] = 1
    }
    
    for i := 1; i < len(dp); i++ {
        for j := 1; j < len(dp[0]); j++ {
            coinVal := coins[i-1]
            am := j
            if coinVal > am {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j] + dp[i][j-coinVal]
            }
        }
    }
    
    return dp[len(dp)-1][len(dp[0])-1]
}