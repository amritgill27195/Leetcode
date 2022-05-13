func maxSubArray(nums []int) int {
    rs := nums[0]
    max := rs
    
    for i := 1; i < len(nums); i++ {
        rs = int(math.Max(float64(rs+nums[i]), float64(nums[i])))
        max = int(math.Max(float64(rs), float64(max)))
    }
    return max
}