func nextPermutation(nums []int)  {
    
    smallestFromRight := -1
    
    for i := len(nums)-2;  i>=0; i-- {
        if nums[i] < nums[i+1] {
            smallestFromRight = i
            break
        }
    }
    if smallestFromRight != -1 {
        // look for immediate number greater than smallestToRight
        smallestVal := nums[smallestFromRight]
        nextGreater := -1

        for i := len(nums)-1; i > smallestFromRight; i-- {
            if nums[i] > smallestVal {
                diff := nums[i] - smallestVal
                if diff > 0 {
                    if nextGreater == -1 {
                        nextGreater = i
                    } else if diff < nums[nextGreater]-smallestVal {
                        nextGreater = i
                    } 
                }   
            }
        }
        
        nums[smallestFromRight], nums[nextGreater] = nums[nextGreater], nums[smallestFromRight]
        
    }
    
    rev(smallestFromRight+1, len(nums)-1, nums)
    
}


func rev(left, right int, nums []int) []int {
    for left < right {
        nums[left], nums[right] = nums[right], nums[left]
        left++
        right--
    }
    return nums
}