func canJump(nums []int) bool {
    destIdx := len(nums)-1
    for i := len(nums)-2; i >= 0 ; i-- {
        if i+nums[i] >= destIdx {
            destIdx = i
        }
    }
    return destIdx==0
}