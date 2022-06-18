func findMiddleIndex(nums []int) int {
    totalSum := 0
    for i := 0; i < len(nums); i++ {
        totalSum += nums[i]
    }
    
    leftSum := 0
    for i := 0; i < len(nums); i++ {
        rightSumExcludingI := totalSum-leftSum-nums[i]
        if rightSumExcludingI == leftSum {
            return i
        }
        leftSum += nums[i]
    }
    return -1
}