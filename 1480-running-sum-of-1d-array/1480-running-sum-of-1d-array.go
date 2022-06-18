func runningSum(nums []int) []int {
    rs := 0
    out := []int{}
    for i := 0; i < len(nums); i++ {
        rs += nums[i]
        out = append(out, rs)
    }
    return out
}