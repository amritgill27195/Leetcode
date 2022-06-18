func pivotIndex(nums []int) int {
    leftSum := []int{nums[0]}
    for i := 1; i < len(nums) ;i++ {
        leftSum = append(leftSum, leftSum[i-1]+nums[i])
    }
    rightSum := make([]int, len(nums))
    rightSum[len(rightSum)-1] = nums[len(nums)-1]
    for i := len(nums)-2; i >=0 ; i-- {
        rightSum[i] = nums[i] + rightSum[i+1]
    }
    
    for i := 0; i < len(nums); i++ {
        if leftSum[i] == rightSum[i] {
            return i
        }
    }
    
    return -1
}