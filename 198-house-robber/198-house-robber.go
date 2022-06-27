func rob(nums []int) int {
    m := len(nums)+1
    n := 2
    dp := make([][]int, m)
    for idx, _ := range dp {
        dp[idx] = make([]int, n)
    }
    for i := len(dp)-2; i >= 0; i-- { // house val
        dp[i][0] = int(math.Max(float64(dp[i+1][0]), float64(dp[i+1][1])))
        dp[i][1] = nums[i] + dp[i+1][0]
    }
    
    return int(math.Max(float64(dp[0][0]),float64(dp[0][1])))
    
}