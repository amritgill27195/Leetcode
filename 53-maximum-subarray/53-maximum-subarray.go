

func maxSubArray(nums []int) int {
    maxSum := int(math.Inf(-1))
    running := 0
    for i := 0; i < len(nums); i++ {
        running += nums[i]
        if nums[i] > running {
            running = nums[i]
        }
        if running > maxSum {
            maxSum = running
        }
    }
    return maxSum
}