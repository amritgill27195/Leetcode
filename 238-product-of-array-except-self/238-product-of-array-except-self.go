func productExceptSelf(nums []int) []int {
    if nums == nil {
        return nil
    }
    left := []int{1}
    for i := 1; i < len(nums); i++ {
        left = append(left, left[i-1] * nums[i-1])
    }
    right := make([]int, len(nums))
    right[len(right)-1] = 1
    for i := len(right)-2; i >= 0; i-- {
        right[i] = right[i+1] * nums[i+1]
    }
    out := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        out[i] = left[i] * right[i]
    }
    return out
}