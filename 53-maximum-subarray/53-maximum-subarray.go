

func maxSubArray(nums []int) int {
    running := 0
    max := math.MinInt64
    for i := 0; i < len(nums); i++ {
        running += nums[i]
        running = int(math.Max(float64(running),float64(nums[i])))
        max = int(math.Max(float64(max), float64(running)))
    }
    return max
}