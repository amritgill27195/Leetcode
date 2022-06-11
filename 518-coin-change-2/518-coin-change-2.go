func change(amount int, coins []int) int {
    n := amount
    dp := make([]int, n+1)
    dp[0] = 1
    
    for i := 0; i < len(coins); i++ {
        for j := 1; j < len(dp);j++ {
            am := j
            coinVal := coins[i]
            if coinVal > am {continue}
            dp[j] = dp[j] + dp[j-coinVal]
        }
    }
    return dp[len(dp)-1]
    
}