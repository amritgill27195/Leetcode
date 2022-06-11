func minCost(costs [][]int) int {
    m := len(costs)
    n := len(costs[0])
    
    dp := make([][]int, m)
    for idx, _ := range dp {
        dp[idx] = make([]int, n)
    }
    
    for j := 0; j < len(dp[0]); j++ {
        dp[m-1][j] = costs[m-1][j]
    }
    
    for i := len(dp)-2; i >= 0; i-- {
        dp[i][0] = costs[i][0] + int(math.Min(float64(dp[i+1][1]),float64(dp[i+1][2])))
        dp[i][1] = costs[i][1] + int(math.Min(float64(dp[i+1][0]),float64(dp[i+1][2])))
        dp[i][2] = costs[i][2] + int(math.Min(float64(dp[i+1][0]),float64(dp[i+1][1])))
    } 
    return int(math.Min(float64(dp[0][0]), math.Min(float64(dp[0][1]), float64(dp[0][2]))))
}