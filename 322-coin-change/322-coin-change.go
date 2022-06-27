func coinChange(coins []int, amount int) int {
    m := len(coins)
    n := amount
    
    dp := make([][]int, m+1)
    for idx, _ := range dp {
        dp[idx] = make([]int, n+1)
    }
    
    for j := 1; j < len(dp[0]); j++ {
        dp[0][j] = amount+1
    }
    
    for i := 1; i < len(dp); i++ {
        for j := 1; j < len(dp[0]); j++ {
            coinVal := coins[i-1]
            am := j
            if coinVal > am {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-coinVal]+1)))
            }
        }
    }
    
    out := dp[m][n]
    if out == amount+1 {return -1}
    return out
    
}