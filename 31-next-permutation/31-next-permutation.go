func nextPermutation(nums []int)  {
    
    
    // find the next smallest number from right
    nextSmaller := -1
    for i := len(nums)-2; i >= 0; i-- {
        if nums[i] < nums[i+1] {
            nextSmaller = i
            break
        }
    }
    
    // find the next greater to nextSmaller
    nextGreater := -1
    if nextSmaller != -1 {
        for i := len(nums)-1; i > nextSmaller; i-- {
            if nums[i] > nums[nextSmaller] {
                if nextGreater == -1 || nums[i] < nums[nextGreater] {
                    nextGreater = i
                }
            }
        }
        nums[nextGreater], nums[nextSmaller] = nums[nextSmaller], nums[nextGreater]
    }
    
    
    left := nextSmaller+1
    right := len(nums)-1
    for left < right {
        nums[left],nums[right] = nums[right], nums[left]
        left++
        right--
    }
}