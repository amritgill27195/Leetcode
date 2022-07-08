func maxProduct(nums []int) int {
    dp := make([][]int, len(nums))
    for idx, _ := range dp {
        dp[idx] = make([]int, 2)
    }
    dp[0] = []int{nums[0], nums[0]}
    
    maxProd := nums[0]
    for i := 1; i < len(nums); i++ {
        dp[i][0] = max(nums[i], max(nums[i] * dp[i-1][0], nums[i] * dp[i-1][1]))
        dp[i][1] = min(nums[i], min(nums[i] * dp[i-1][0], nums[i] * dp[i-1][1]))
        maxProd = max(maxProd, max(dp[i][0], dp[i][1]))
    }
    
    return maxProd
}
func max(x, y int) int {
    if x > y {return x}
    return y
}
func min(x, y int) int {
    if x < y {return x}
    return y
}