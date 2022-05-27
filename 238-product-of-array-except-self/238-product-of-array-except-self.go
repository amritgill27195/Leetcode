func productExceptSelf(nums []int) []int {
    lProd := make([]int, len(nums))
    lProd[0] = 1
    for i := 1; i < len(nums); i++ {
        prev := nums[i-1]
        prevProd := lProd[i-1]
        lProd[i] = prev * prevProd
    }
    
    rProd := make([]int, len(nums))
    rProd[len(nums)-1] = 1
    for i := len(nums)-2; i >= 0; i-- {
        prev := nums[i+1]
        prevProd := rProd[i+1]
        rProd[i] = prev*prevProd
    }
    
    out := []int{}
    for i := 0; i < len(lProd); i++ {
        out = append(out, lProd[i] * rProd[i])
    }
    return out
}