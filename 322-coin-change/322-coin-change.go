func coinChange(coins []int, amount int) int {
    m := len(coins)
    n := amount
    
    dp := make([]int, n+1)
    for j := 1; j < len(dp); j++ {
        dp[j] = amount+1
    }
    
    for i := 1; i <= m; i++ {
        for j := 1; j <= amount; j++ {
            coinVal := coins[i-1]
            am := j
            if coinVal <= am {
                dp[j] = int(math.Min(float64(dp[j]), float64(dp[j-coinVal]+1)))
            }
        }
    }
    
    out := dp[n]
    if out == amount+1 {return -1}
    return out
    
}