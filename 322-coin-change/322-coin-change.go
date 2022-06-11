func coinChange(coins []int, amount int) int {
    n := amount
    dp := make([]int, n+1)
    for i := 1; i < len(dp); i++ {
        dp[i] = amount+1
    } 
    for i := 0; i < len(coins); i++ {
        for j := 0; j < len(dp); j++ {
            coinVal := coins[i]
            am := j
            if coinVal > am {
                continue
            }
            dp[j] = int(math.Min(float64(dp[j]),float64(dp[j-coinVal]+1)))
        }
    }
    result := dp[len(dp)-1]
    if result == amount+1 {
        return -1
    }
    return result
}