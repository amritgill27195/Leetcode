func productExceptSelf(nums []int) []int {
    left := make([]int, len(nums))
    left[0] = 1
    for i := 1; i < len(nums); i++ {
        prevNum := nums[i-1]
        prevProd := left[i-1]
        left[i] = prevNum * prevProd
    }
    right := make([]int, len(nums))
    right[len(right)-1] = 1
    for i := len(nums)-2; i >= 0; i-- {
        prevNum := nums[i+1]
        prevProd := right[i+1]
        right[i] = prevNum * prevProd
    }
    
    out := []int{}
    for i := 0; i < len(nums); i++ {
        out = append(out, left[i] * right[i])
    }
    return out
}