func productExceptSelf(nums []int) []int {
    out := make([]int, len(nums))
    out[0] = 1
    for i := 1; i < len(nums); i++ {
        prevNum := nums[i-1]
        prevProd := out[i-1]
        out[i] = prevNum * prevProd
    }
    
    prp := 1
    for i := len(nums)-2; i >= 0; i-- {
        prevNum := nums[i+1]
        prp = prp * prevNum
        out[i] *= prp
    }
    return out
}