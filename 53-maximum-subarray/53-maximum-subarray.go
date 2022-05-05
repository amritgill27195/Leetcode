

func maxSubArray(nums []int) int {
    running := 0
    max := math.MinInt64
    for i := 0; i < len(nums); i++ {
        running += nums[i]
        running = getMax(running, nums[i])
        max = getMax(running, max)
    }
    return max
}

func getMax(n,m int) int {
    if n > m {return n}
    return m
}