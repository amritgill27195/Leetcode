func numberOfArithmeticSlices(nums []int) int {
    dp := make([]int, len(nums))
    n := len(nums)
    sum := 0
    
    for i := n-3; i >= 0; i-- {
        if nums[i+1] - nums[i] == nums[i+2] - nums[i+1] {
            dp[i] = dp[i+1] + 1
            sum += dp[i]
        }
    }
    return sum
}