func numberOfArithmeticSlices(nums []int) int {
    dp := make([]int, len(nums))
    count := 0
    for i := 2; i < len(nums); i++ {
        if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
            dp[i] = dp[i-1]+1
            count += dp[i]
        }
    }
    return count
}