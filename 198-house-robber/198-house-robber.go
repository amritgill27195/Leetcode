func rob(nums []int) int {
    n := 2
    dp := make([]int, n)

    for i := len(nums)-1; i >= 0; i-- {
        houseVal := nums[i]
        zeroCase := dp[0]
        oneCase := dp[1]
        
        dp[0] = int(math.Max(float64(zeroCase), float64(oneCase)))
        dp[1] = zeroCase + houseVal
        
    }
    
    return int(math.Max(float64(dp[0]),float64(dp[1])))
    
}