func maxSubArray(nums []int) int {
    max := nums[0]
    rs := nums[0]
    
    for i := 1; i < len(nums); i++ {
        rs = int(math.Max(float64(nums[i]), float64(rs+nums[i])))
        max = int(math.Max(float64(max),float64(rs))) 
    }
    return max
}