func nextPermutation(nums []int)  {
    breachIdx := -1
    for i := len(nums)-2; i >= 0; i-- {
        if nums[i] < nums[i+1] {
            breachIdx = i
            break
        }
    }
    
    if breachIdx != -1 {
        nextGreaterOfBreachIdx := -1
        
        for i := len(nums)-1; i > breachIdx; i-- {
            if nums[i] > nums[breachIdx] {
                if nextGreaterOfBreachIdx == -1 {
                    nextGreaterOfBreachIdx = i
                } else {
                    if nums[i] - nums[breachIdx] < nums[nextGreaterOfBreachIdx]-nums[breachIdx] {
                        nextGreaterOfBreachIdx = i
                    }
                }
            }
        }
        
        nums[breachIdx], nums[nextGreaterOfBreachIdx] = nums[nextGreaterOfBreachIdx], nums[breachIdx]
    }
    
    nums = reverse(breachIdx+1, len(nums)-1, nums)
}


func reverse(left, right int, nums []int) []int {
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
    return nums
}